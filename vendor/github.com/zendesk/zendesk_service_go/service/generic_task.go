package service

import "context"

// GenericTask is an adapter to use functions as service.Process implementations
type GenericTask func(context.Context) error

// Start will execute the generic task
// tasks are expected to respect context cancellation and shut down appropriately
func (task GenericTask) Start(ctx context.Context) error {
	return task(ctx)
}

var _ Process = GenericTask(nil)
