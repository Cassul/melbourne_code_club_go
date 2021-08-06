package errorcollector

import (
	"context"
	"sync"
)

type errorCollectorContextKey string

const contextKey errorCollectorContextKey = "error_collector_context"

// InjectCollector will inject an error collector into the context
//
// if there is already one in the context, then this is a no-op
func InjectCollector(ctx context.Context) context.Context {
	collector := fromContext(ctx)
	if collector != nil {
		return ctx
	}

	ctx = context.WithValue(ctx, contextKey, &requestErrorCollector{})

	return ctx
}

// RecordError will record an error into the current context's error collector
//
// if no error collector has been added to the context, then this is a no-op
func RecordError(ctx context.Context, err error) {
	collector := fromContext(ctx)
	if collector == nil {
		return
	}

	collector.setError(err)
}

// ErrorFromContext returns any error that has been attached to the context's error collector
//
// If no error collector has been attached to the context, then this is a no-op
func ErrorFromContext(ctx context.Context) error {
	collector := fromContext(ctx)
	if collector == nil {
		return nil
	}

	return collector.getError()
}

// requestErrorCollector is a type that is injected to the ctx.Context at the beginning of a request
// and will record any error that occurs within the request
//
// This isn't intended to be used directly in application code, but to be used by things like
// logging / error reporting libraries, and other middleware
type requestErrorCollector struct {
	errLock sync.Mutex
	err     error
}

func (collector *requestErrorCollector) setError(err error) {
	collector.errLock.Lock()
	defer collector.errLock.Unlock()

	collector.err = err
}

func (collector *requestErrorCollector) getError() error {
	collector.errLock.Lock()
	defer collector.errLock.Unlock()

	return collector.err
}

func fromContext(ctx context.Context) *requestErrorCollector {
	collector, ok := ctx.Value(contextKey).(*requestErrorCollector)
	if ok {
		return collector
	}

	return nil
}
