package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	logger "github.com/zendesk/zendesk_logging_go"
	"golang.org/x/sync/errgroup"
)

const defaultShutdownMinutes = 3

var ErrShutdownTimeout = fmt.Errorf("application shutdown timed out")

// Application contains one or more Processes, and will run until they complete or fail
type Application struct {
	subprocesses map[string]Process
	processLock  sync.Mutex

	// ShutdownTimeout configures how long we have for an application to shut down cleanly
	// when one subprocess exits before we pull the plug.
	// In the future we can configure this with functional options, but the default of 3 minutes
	// should be appropriate in most cases - so just make this field public for now.
	ShutdownTimeout time.Duration
}

// NewApplication initializes a new application
func NewApplication() *Application {
	return &Application{
		subprocesses:    map[string]Process{},
		ShutdownTimeout: defaultShutdownMinutes * time.Minute,
	}
}

// AddSubprocess adds a named subprocess to the application
// reusing a name wil replace a given subprocess
func (app *Application) AddSubprocess(name string, process Process) {
	app.processLock.Lock()
	defer app.processLock.Unlock()

	app.subprocesses[name] = process
}

// Subprocess returns a Process from the application collection by name
func (app *Application) Subprocess(name string) Process {
	app.processLock.Lock()
	defer app.processLock.Unlock()

	return app.subprocesses[name]
}

// Run starts each of the registered subprocesses
func (app *Application) Run() error {
	return app.RunWithContext(context.Background())
}

// RunWithContext starts each of the registered subprocesses with a given Context
func (app *Application) RunWithContext(ctx context.Context) error {
	app.processLock.Lock()
	defer app.processLock.Unlock()

	if len(app.subprocesses) == 0 {
		return fmt.Errorf("no processes registered")
	}

	// appCtx gets threaded to all subprocesses
	// subprocesses are expected to gracefully shutdown when context is cancelled
	appCtx, cancel := context.WithCancel(ctx)

	// signal handling
	signalChan, signalShutdown := initSignalHandler()
	defer signalShutdown()

	// errgroup's context will be cancelled if any of the child goroutines return error
	// also _explicitly_ cancel appCtx (via `cancel()`) if any process exits with nil error
	// in that case, ensure the app actually exits - thus either *all* subprocesses are running, or we shutdown
	group, errgroupCtx := errgroup.WithContext(appCtx)
	group.Go(func() error {
		select {
		case <-appCtx.Done():
			// a subprocess has gracefully exited - call to cancel() made
		case caughtSig := <-signalChan:
			// signal received by process, shut down subprocesses by cancelling context
			logger.Infof(appCtx, "%s signal received, shutting down", caughtSig)
			cancel()
		}

		return nil
	})

	for name, process := range app.subprocesses {
		// re-alias the loop variables to prevent them getting clobbered on next iteration
		name := name
		process := process

		group.Go(func() error {
			ctx := logger.WithField(errgroupCtx, "subprocess", name)
			if err := process.Start(ctx); err != nil {
				// subprocess error
				logger.FromContext(ctx).WithError(err).Error("subprocess error, shutting down")
				return err
			} else {
				logger.FromContext(ctx).Infof("subprocess exited gracefully, shutting down")
			}

			// note: ok to cancel appCtx multiple times - is safely a noop for subsequent calls
			cancel()
			return nil
		})
	}

	// This is guaranteed to fire. Either one of errgroup's subprocesses return an error, in which case
	// the errgroup cancels context, or they exit successfully, in which case cancel() is called in each
	// process.Start(ctx) == nil check (or the signal handler), and thus cancellation cascades down to errgroupCtx
	<-errgroupCtx.Done()
	return app.waitForFinish(group, appCtx)
}

// waitForFinish will await for given application error group to finish execution
func (app *Application) waitForFinish(group *errgroup.Group, appCtx context.Context) error {
	// We want to be graceful, but it's dangerous to wait forever because of a badly coded subprocess
	// e.g. there might be a debug server or such running, keeping our process alive, but not serving any traffic.
	// So as an escape hatch, force-exit without cleanup if we exceed our timeout.
	errGroupDone := make(chan error, 1)
	go func() {
		errGroupDone <- group.Wait()
		close(errGroupDone)
	}()

	// await for error group done, or timeout exceeded
	select {
	case err := <-errGroupDone:
		return err
	case <-time.After(app.ShutdownTimeout):
		logger.FromContext(appCtx).Error("application timed out shutting down. force-exiting...")
		return ErrShutdownTimeout
	}
}

func initSignalHandler() (chan os.Signal, func()) {
	// make a channel for signal notifications and config signals to listen for
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// return channel and a shutdown function to be called by defer
	return signalChan, func() {
		signal.Stop(signalChan)
		close(signalChan)
	}
}
