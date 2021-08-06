package logger

import (
	"net/http"
	"time"

	"github.com/zendesk/zendesk_langsupport_go/errorcollector"
	"github.com/zendesk/zendesk_langsupport_go/responsewriter"
)

// HTTPMiddleware is a HTTP middleware function for logging requests
func HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		wrappedResponse := responsewriter.NewStatusCapturingResponseWriter(response)

		ctx := WithFields(request.Context(), Fields{
			"request.method": request.Method,
			"request.path":   request.URL.Path,
			// This looks strange (request_id instead of request.id) but the observability standard
			// says that's what it's called.
			"request_id": request.Header.Get("X-Request-Id"),
		})
		ctx = errorcollector.InjectCollector(ctx)

		*request = *request.WithContext(ctx)
		next.ServeHTTP(wrappedResponse, request)

		fields := Fields{
			"response.status":   wrappedResponse.GetResponseStatusCode(),
			"response.duration": time.Since(startTime),
		}

		err := errorcollector.ErrorFromContext(ctx)
		if err != nil {
			fields["error"] = err
		}

		ctx = WithFields(request.Context(), fields)

		// Any client or server error gets reported as error
		if wrappedResponse.GetResponseStatusCode() >= http.StatusInternalServerError {
			Errorf(ctx, "%s completed", request.URL.Path)
		} else {
			Infof(ctx, "%s completed", request.URL.Path)
		}
	})
}
