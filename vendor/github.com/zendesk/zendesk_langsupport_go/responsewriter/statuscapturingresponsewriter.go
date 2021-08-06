package responsewriter

import "net/http"

type StatusCapturingResponseWriter interface {
	http.ResponseWriter
	GetResponseStatusCode() int
}

// NewStatusCapturingResponseWriter wraps the provided response writer in a wrapper that captures the status
// for you, so you can react to it in middleware.
func NewStatusCapturingResponseWriter(inner http.ResponseWriter) StatusCapturingResponseWriter {
	// The order of this switch is important, because the upper interfaces are supersets of the lower ones.
	switch innerTyped := inner.(type) {
	case StatusCapturingResponseWriter:
		// Don't re-wrap if some middleware already wrapped.
		return innerTyped
	case flusherPusherHijackerWriter:
		return &flusherPusherHijackerWriterImpl{innerTyped, 0}
	case flusherPusherWriter:
		return &flusherPusherWriterImpl{innerTyped, 0}
	case pusherHijackerWriter:
		return &pusherHijackerWriterImpl{innerTyped, 0}
	case flusherHijackerWriter:
		return &flusherHijackerWriterImpl{innerTyped, 0}
	case flusherWriter:
		return &flusherWriterImpl{innerTyped, 0}
	case pusherWriter:
		return &pusherWriterImpl{innerTyped, 0}
	case hijackerWriter:
		return &hijackerWriterImpl{innerTyped, 0}
	default:
		return &bareWriterImpl{innerTyped, 0}
	}
}

// We need an interface type for every possible combination of HTTP interfaces (except http.ResponseWriter is
// of course in all of them).
type flusherWriter interface {
	http.Flusher
	http.ResponseWriter
}
type pusherWriter interface {
	http.Pusher
	http.ResponseWriter
}
type hijackerWriter interface {
	http.Hijacker
	http.ResponseWriter
}
type flusherPusherWriter interface {
	http.Flusher
	http.Pusher
	http.ResponseWriter
}
type pusherHijackerWriter interface {
	http.Pusher
	http.Hijacker
	http.ResponseWriter
}
type flusherHijackerWriter interface {
	http.Flusher
	http.Hijacker
	http.ResponseWriter
}
type flusherPusherHijackerWriter interface {
	http.Flusher
	http.Pusher
	http.Hijacker
	http.ResponseWriter
}

// bareWriterImpl implements http.ResponseWriter, but none of the other optional HTTP interfaces
type bareWriterImpl struct {
	http.ResponseWriter
	statusCode int
}

func (i *bareWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.ResponseWriter.WriteHeader(statusCode)
}

func (i *bareWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.ResponseWriter.Write(data)
}

func (i *bareWriterImpl) GetResponseStatusCode() int {
	return i.statusCode
}

// Now for all of the possible HTTP interfaces defined above, define a struct that implements them.
type flusherWriterImpl struct {
	flusherWriter
	statusCode int
}

func (i *flusherWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.flusherWriter.WriteHeader(statusCode)
}
func (i *flusherWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.flusherWriter.Write(data)
}
func (i *flusherWriterImpl) GetResponseStatusCode() int { return i.statusCode }

type pusherWriterImpl struct {
	pusherWriter
	statusCode int
}

func (i *pusherWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.pusherWriter.WriteHeader(statusCode)
}
func (i *pusherWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.pusherWriter.Write(data)
}
func (i *pusherWriterImpl) GetResponseStatusCode() int { return i.statusCode }

type hijackerWriterImpl struct {
	hijackerWriter
	statusCode int
}

func (i *hijackerWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.hijackerWriter.WriteHeader(statusCode)
}
func (i *hijackerWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.hijackerWriter.Write(data)
}
func (i *hijackerWriterImpl) GetResponseStatusCode() int { return i.statusCode }

type flusherPusherWriterImpl struct {
	flusherPusherWriter
	statusCode int
}

func (i *flusherPusherWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.flusherPusherWriter.WriteHeader(statusCode)
}
func (i *flusherPusherWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.flusherPusherWriter.Write(data)
}
func (i *flusherPusherWriterImpl) GetResponseStatusCode() int { return i.statusCode }

type pusherHijackerWriterImpl struct {
	pusherHijackerWriter
	statusCode int
}

func (i *pusherHijackerWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.pusherHijackerWriter.WriteHeader(statusCode)
}
func (i *pusherHijackerWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.pusherHijackerWriter.Write(data)
}
func (i *pusherHijackerWriterImpl) GetResponseStatusCode() int { return i.statusCode }

type flusherHijackerWriterImpl struct {
	flusherHijackerWriter
	statusCode int
}

func (i *flusherHijackerWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.flusherHijackerWriter.WriteHeader(statusCode)
}
func (i *flusherHijackerWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.flusherHijackerWriter.Write(data)
}
func (i *flusherHijackerWriterImpl) GetResponseStatusCode() int { return i.statusCode }

type flusherPusherHijackerWriterImpl struct {
	flusherPusherHijackerWriter
	statusCode int
}

func (i *flusherPusherHijackerWriterImpl) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.flusherPusherHijackerWriter.WriteHeader(statusCode)
}
func (i *flusherPusherHijackerWriterImpl) Write(data []byte) (int, error) {
	if i.statusCode == 0 {
		i.statusCode = 200
	}
	return i.flusherPusherHijackerWriter.Write(data)
}
func (i *flusherPusherHijackerWriterImpl) GetResponseStatusCode() int { return i.statusCode }
