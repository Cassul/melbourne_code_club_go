package service

import (
	"context"
	"time"
)

// Process is an abstraction for any sub-process
type Process interface {
	// Start starts the sub-process which should gracefully shutdown when given context is cancelled
	Start(context.Context) error
}

// ProcessTiming wraps a given Process and tracks start and stop time for the process
type ProcessTiming struct {
	Process

	startTime time.Time
	endTime   time.Time
}

// NewProcessTiming creates a new ProcessTiming wrapped process
func NewProcessTiming(process Process) *ProcessTiming {
	return &ProcessTiming{Process: process}
}

// Start commences the wrapped process, remembering both start and stop times
func (pt *ProcessTiming) Start(ctx context.Context) error {
	pt.startTime = time.Now()
	err := pt.Process.Start(ctx)
	pt.endTime = time.Now()

	return err
}

// Runtime returns the total running duration for the wrapped process
func (pt *ProcessTiming) Runtime() time.Duration {
	if pt.startTime.IsZero() || pt.endTime.IsZero() {
		// process either not started/finished
		return time.Duration(-1)
	}

	return pt.endTime.Sub(pt.startTime)
}
