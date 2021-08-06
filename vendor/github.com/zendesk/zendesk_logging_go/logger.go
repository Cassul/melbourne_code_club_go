package logger

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// FormatText constant, used to set the log output format to text
	FormatText = "text"
	// FormatJSON constant, used to set the log output format to json
	FormatJSON = "json"

	// LevelInfo constant, used to set the log level to info
	LevelInfo = "info"
	// LevelWarn constant, used to set the log level to warn
	LevelWarn = "warn"
	// LevelError constant, used to set the log level to error
	LevelError = "error"
	// LevelDebug constant, used to set the log level to debug
	LevelDebug = "debug"
)

// LogEntry type, an alias for a logrus.Entry
type LogEntry = logrus.Entry

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

func (fields Fields) empty() bool {
	return len(fields) == 0
}

var globalConfig Config

// Setup takes a list of Option setters and configures the logger
func Setup(setters ...Option) error {
	config := defaultConfig

	for _, setter := range setters {
		setter(&config)
	}

	setFormat(config)

	err := setLevel(config.Level)
	if err != nil {
		return err
	}

	setOutput(config.Output)
	setDefaultFields(config.DefaultFields)
	setFieldExtractorFuncs(config.FieldExtractorFuncs...)
	return nil
}

// SetupWithConfig takes a Config object loaded using zendesk_config_go to configure the logger
func SetupWithConfig(config Config, setters ...Option) error {
	options := append(optionsFromConfig(config), setters...)
	return Setup(options...)
}

// FromContext returns a new *logrus.Entry with fields populated from the context
func FromContext(ctx context.Context) *LogEntry {
	ctxWithDefaultFields := WithFields(ctx, globalConfig.DefaultFields)
	return extractLoggerDataWithFieldExtractors(ctxWithDefaultFields).takeReferenceAsLogrusEntry()
}

// Info logs a message at the info level
func Info(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Info(args...)
}

// WithField returns a new child context.Context with a given logging field added
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	data := extractLoggerDataWithFieldExtractors(ctx)
	data.fields[key] = value
	return context.WithValue(ctx, loggerCtxKey, data)
}

// WithFields returns a new child context.Context with a given logging fields added
func WithFields(ctx context.Context, fields Fields) context.Context {
	if len(fields) == 0 {
		return ctx
	}

	data := extractLoggerDataWithFieldExtractors(ctx)
	for key, value := range fields {
		data.fields[key] = value
	}
	return context.WithValue(ctx, loggerCtxKey, data)
}

// ExtractFields returns a shallow copy of the logging fields from a given context
func ExtractFields(ctx context.Context) Fields {
	return extractLoggerDataWithFieldExtractors(ctx).fields
}

// Infof logs a formatted message at the info level
func Infof(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Infof(format, args...)
}

// Warn logs a message at the warn level
func Warn(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Warn(args...)
}

// Warnf logs a formatted message at the warn level
func Warnf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Warnf(format, args...)
}

// Error logs a message at the error level
func Error(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Error(args...)
}

// Errorf logs a formatted message at the error level
func Errorf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Errorf(format, args...)
}

// Debug logs a message at the debug level
func Debug(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Debug(args...)
}

// Debugf logs a formatted message at the debug level
func Debugf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Debugf(format, args...)
}

func setLevel(level string) error {
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("could not set log level: %+v", err)
	}

	logrus.SetLevel(parsedLevel)
	globalConfig.Level = level
	return nil
}

func setFormat(config Config) {
	var formatter logrus.Formatter
	switch config.Format {
	case FormatJSON:
		formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  "message",
				logrus.FieldKeyTime: "@timestamp",
			},
			PrettyPrint: config.PrettyPrintJSON,
		}
	case FormatText:
		fallthrough
	default:
		formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "3:04:05PM",
			ForceColors:     true,
		}
	}

	logrus.SetFormatter(formatter)
	globalConfig.Format = config.Format
}

func setOutput(out io.Writer) {
	logrus.SetOutput(out)
	globalConfig.Output = out
}

func setDefaultFields(fields Fields) {
	globalConfig.DefaultFields = fields
}

func setFieldExtractorFuncs(funcs ...FieldExtractorFunc) {
	globalConfig.FieldExtractorFuncs = funcs
}
