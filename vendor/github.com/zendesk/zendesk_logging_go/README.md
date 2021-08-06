# zendesk_logging_go
![Test](https://github.com/zendesk/zendesk_logging_go/workflows/Test/badge.svg)

Handle JSON logging for Zendesk Golang applications.

This is a wrapper over [logrus](https://github.com/sirupsen/logrus) that assists in the implementation of the [Zendesk logging standard](https://techmenu.zende.sk/standards/observability/).

 > *Note:*
 > In order to fully comply with the Zendesk Logging Standard, you will need to implement [a custom `FieldExtractor` function](https://github.com/zendesk/zendesk_logging_go#configuring-dynamic-default-fields-via-a-fieldextractor-function) to retrieve the following fields from the request context and add them to your log messages:
 > * `account.id`
 > * `account.subdomain`
 > * `request_id`
 > * `dd.trace_id`
 > * `dd.span_id`
 >
 > You will also need to add a `pod` field. This can be retrieved from an environment variable and added to log messages using the [`logger.DefaultFields()` option](https://github.com/zendesk/zendesk_logging_go#custom-options).
 >

### TODO

- [ ] Custom formatters
- [ ] Document configuration examples to make it clear how to conform to the logging standard

## Installation
Follow the guide at https://zendesk.atlassian.net/wiki/spaces/ENG/pages/779198045/Private+Go+Modules+at+Zendesk

## Environment variables

Configuring the logger the recommended way with [zendesk_config_go](https://github.com/zendesk/zendesk_config_go) assumes the following environment variables are provided to your application. All the variables have default values.

| Key          | Description           | Default  |
| ------------ | --------------------- | -------- |
| `LOG_LEVEL`  | Minimum level of logs | `"info"` |
| `LOG_FORMAT` | Output format of logs | `"json"` |

## Usage

### Basic Logging

```go
import (
  "context"
  "fmt"

  logger "github.com/zendesk/zendesk_logging_go"
)

func main() {
  err := logger.Setup()
  if err != nil {
    fmt.Println(err)
  }

  ctx := context.Background()
  logger.Info(ctx, "this is a info log")
  logger.Warn(ctx, "this is a warn log")
  logger.Error(ctx, "this is a error log")
}
```

### With zendesk_config_go

If you're using [zendesk_config_go](https://github.com/zendesk/zendesk_config_go) to load your applications configuration. The same `Config` struct can be used to configure the logger.

The following environment variables are used:
* `LOG_LEVEL` for the level (defaults to logger.LevelInfo)
* `LOG_FORMAT` for the format (defaults to logger.FormatJSON)

```go
import (
  "context"
  "fmt"
  configuration "github.com/zendesk/zendesk_config_go"
  logger "github.com/zendesk/zendesk_logging_go"
)

type Config struct {
  // application config properties
  // ...

  Logging logger.Config
}

func main() {
  config := Config{}
  err := configuration.Load(&config)
  if err != nil {
    panic(err)
  }


  err := logger.SetupWithConfig(config.Logging)
  // You can also can override what gets loaded into config
  // err := logger.SetupWithConfig(config, logger.Level(logger.LevelWarn), ...)
  if err != nil {
    fmt.Println(err)
  }

  ctx := context.Background()
  logger.Info(ctx, "this is a info log")
  logger.Warn(ctx, "this is a warn log")
  logger.Error(ctx, "this is a error log")
}
```

### Printf Style Logging

```go
import (
  "context"
  "fmt"

  logger "github.com/zendesk/zendesk_logging_go"
)

func main() {
  err := logger.Setup()
  if err != nil {
    fmt.Println(err)
  }

  ctx := context.Background()
  logger.Infof(ctx, "this is a info log with formatting: %d %d", 1, 2)
  logger.Warnf(ctx, "this is a warn log with formatting: %d %d", 1, 2)
  logger.Errorf(ctx, "this is a error log with formatting: %d %d", 1, 2)
}
```

### Custom Options

```go
import (
  "context"
  "fmt"
  "os"

  logger "github.com/zendesk/zendesk_logging_go"
)

func main() {
  // These are all the default options
  err := logger.Setup(
    logger.Level(logger.LevelInfo),   // Can be logger.LevelDebug, logger.LevelInfo, logger.LevelWarn, logger.LevelError
    logger.Format(logger.FormatJSON), // Can be logger.FormatText or logger.FormatJSON
    logger.PrettyPrintJSON(false),    // Can be enabled in dev for multi-line, indented json output.
    logger.Output(os.Stdout),         // Can be any type that implements the io.Writer interface
    logger.DefaultFields(Fields{}),   // Default Fields for every log
  })
  if err != nil {
    fmt.Println(err)
  }

  ...
}
```

### Logging with Fields
Context fields can be automatically added to the logs when creating a log entry by generating a new context with the `logger.WithField(...)` or `logger.WithFields(...)` functions. You can retrieve these fields at any point via `logger.FromContext(ctx)`.

```go
import (
  "context"
  "fmt"

  logger "github.com/zendesk/zendesk_logging_go"
)

func main() {
  err := logger.Setup()
  if err != nil {
    fmt.Println(err)
  }

  // With a single field
  ctx := logger.WithField(context.Background(), "foo", "bar")
  logger.Info(ctx, "info log with field")
  // Will log something like:
  // {"@timestamp": "2019-10-13T13:16:38.662Z", "level": "info", "foo": "bar", "message": "info log with field"}

  // With multiple fields
  ctx = logger.WithFields(context.Background(), logger.Fields{
    "foo": "bar",
    "num": 10,
  })
  logger.Info(ctx, "info log with fields")
  // Will log something like:
  // {"@timestamp": "2019-10-13T13:16:38.662Z", "level": "info", "foo": "bar", "num": 10, "message": "info log with fields"}
}
```

Note: When running applications in Kubernetes, fluentd is configured to automatically add fields to your log output, including @timestamp and application (via metadata.labels.project of the pod). View [this page](https://techmenu.zende.sk/standards/observability/) for more information.

### Configuring Dynamic Default Fields via a `FieldExtractor` function
A `FieldExtractor` function can be configured during the logger setup. It is called whenever fields are added to a logline. This option is useful if you require default fields to be added to every logline (e.g. fields that appear in the context like `trace_id` or `account_id`).

zendesk_logging_go comes built in with two extractors in the `extractors/` directory that should be useful for most apps.

* `AccountIDFieldExtractor`, which is used to read the account ID out of the context. You can feed in the account ID into context by using `ctx = NewContextWithAccountIDString(ctx, acctId)` in your app. Alternatively, you can automagically inject the account ID in GRPC server handler contexts by using the `AccountIDUnaryInterceptor` middleware, which calls the `GetAccountId()` method on your protobufs, if they have one.
* `DataDogFieldExtractor`, which extracts Datadog trace IDs from Datadog APM spans in contexts.

To use them, do something like the following:

```go
import (
  "context"
  "fmt"
  "strconv"

  "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

  logger "github.com/zendesk/zendesk_logging_go"
  "github.com/zendesk/zendesk_logging_go/extractors"
)

func main() {
  err := logger.Setup(
    // multiple field extractors will be processed in order, with subsequent extractors taking precedence in the case of duplicate keys
    logger.FieldExtractorFuncs(extractors.AccountIDFieldExtractor, extractors.DataDogFieldExtractor)
  )
  if err != nil {
    fmt.Println(err)
  }

  // With a single field
  logger.Info(ctx, "info log with field")
  // Will log something like:
  // {"@timestamp": "2019-10-13T13:16:38.662Z", "level": "info", "message": "info log", "dd.trace_id": "12345", "dd.span_id": "67890"}

}

```

### Logging with an Error

```go
import (
  "context"
  "fmt"

  logger "github.com/zendesk/zendesk_logging_go"
)

func main() {
  err := logger.Setup()
  if err != nil {
    fmt.Println(err)
  }

  err = fmt.Errorf("some error")
  if err != nil {
    logLine := logger.FromContext(ctx).WithError(err)
    logLine.Errorf("A Zendesk error occurred: %s", err.Error())
  }
  // Will log something like:
  // {"@timestamp": "2019-10-13T13:16:38.662Z", "level": "error", "error": "some error", "message": "A Zendesk error occurred: some error"}
}
```

### HTTP Middleware

```go
import (
  "fmt"
  "net/http"

  logger "github.com/zendesk/zendesk_logging_go"
)

func main() {
  err := logger.Setup()
  if err != nil {
    fmt.Println(err)
  }

  // Will log request/response details (path, status, method, duration) at INFO level
  http.Handle("/some/path", logger.HTTPMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
  })))
}
```

### gRPC Middleware

zendesk_logging_go provides Unary and Streaming gRPC middleware for both servers and clients. The `UnaryServerMiddleware` and `StreamServerMiddleware` include both gRPC request context decoration and logging interceptors.

```go
import (
  "context"
  "fmt"

  logger "github.com/zendesk/zendesk_logging_go"
  grpclogger "github.com/zendesk/zendesk_logging_go/grpc"
)

func main() {
  err := logger.Configure()
  if err != nil {
    fmt.Println(err)
  }

  logEntry := logger.FromContext(context.Background())

  server := grpc.NewServer(
    grpc_middleware.WithUnaryServerChain(
      grpclogger.UnaryServerInterceptor(logEntry),
      // other interceptors
    ),
    grpc_middleware.WithStreamServerChain(
      grpclogger.StreamServerInterceptor(logEntry),
      // other interceptors
    ),
  )

  // Or you can customize the middleware with options
  server = grpc.NewServer(
    grpc_middleware.WithUnaryServerChain(
      grpclogger.UnaryServerInterceptor(logEntry, grpclogger.FieldExtractor(grpclogger.CodeGenRequestFieldExtractor)),
    ),
    grpc_middleware.WithStreamServerChain(
      grpclogger.StreamServerInterceptor(logEntry, grpclogger.FieldExtractor(grpclogger.TagBasedRequestFieldExtractor("log_field"))),
    ),
  )
}
```

#### Account ID logging

You can automatically log `account.id` with the `AccountIDFieldExtractor` and `AccountIDUnaryInterceptor`. This assumes that `account_id` is a `*wrappers.StringValue`, `*wrappers.Int64Wrapper`, `*wrappers.Int32Value`, or `int64`.
Other account_id types can be added by adding another interceptor of the correct type.

```go
  server := grpc.NewServer(
    grpc_middleware.WithUnaryServerChain(
      grpclogger.AccountIDUnaryInterceptor(),
      // other interceptors
    ),
  )
```
#### User ID logging

You can automatically log `user.id` with the `UserIDFieldExtractor` and `UserIDUnaryInterceptor`. This assumes that `user_id` is a `*wrappers.StringValue`, `string`, `*wrappers.Int64Wrapper` or `int64`.
Other user_id types can be added by adding another interceptor of the correct type.

```go
  server := grpc.NewServer(
    grpc_middleware.WithUnaryServerChain(
      grpclogger.UserIDUnaryInterceptor(),
      // other interceptors
    ),
  )
```

#### Subdomain logging

You can automatically log `subdomain` with the `SubdomainFieldExtractor` and `SubdomainUnaryInterceptor`. This assumes that `subdomain` is a `*wrappers.StringValue` or `string`.

```go
  server := grpc.NewServer(
    grpc_middleware.WithUnaryServerChain(
      grpclogger.SubdomainUnaryInterceptor(),
      // other interceptors
    ),
  )
```


## Logger testing

The `loggertest` package provides some helpers to test logging inside your project.

```go
import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zendesk/zendesk_logging_go/loggertest"
)

func LogAction(ctx context.Context) {
  if err := doSomething(); err != nil { // this returns "doSomethingError" error
    logger.FromContext(ctx).Error(
      fmt.Sprintf("Something went wrong: %s", err.String()),
    )
  }
}

func LogActionWithError(t *testing.T) {
  cf := loggertest.CaptureLogrusLogs(t)

  LogAction(context.Background())

  assert.Equal(t, "Something went wrong: doSomethingError", cf.LastEntry().Message)
}
```

You can have a look at [capturing_formatter_test.go](https://github.com/zendesk/zendesk_logging_go/blob/master/loggertest/capturing_formatter_test.go) for more examples.

## Contributing

Clone the repository into your Zendesk code path e.g `~/Code/zendesk/`.

```shell
cd ~/Code/zendesk
git clone git@github.com:zendesk/zendesk_logging_go.git
cd zendesk_logging_go
make install_devtools
```

**Note: This library uses [Go modules](https://github.com/golang/go/wiki/Modules#go-111-modules), so it does not need to be cloned into `$GOPATH`.**

### Testing

```shell
make test
```

### Linting

```shell
make lint
```
