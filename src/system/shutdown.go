package system

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const MaxWaitDuration = 5 * time.Second

type componentSignal struct {
	doneChan <-chan error
	done     bool
}

type GracefulShutdown struct {
	shutdown   func()
	components map[string]*componentSignal
}

func NewGracefulShutdown(ctx context.Context) (*GracefulShutdown, context.Context) {
	appCtx, cancel := context.WithCancel(ctx)

	return &GracefulShutdown{
		shutdown:   cancel,
		components: map[string]*componentSignal{},
	}, appCtx
}

func (gs *GracefulShutdown) RegisterComponent(name string, doneChan <-chan error) {
	gs.components[name] = &componentSignal{
		doneChan: doneChan,
		done:     false,
	}
}

func (gs *GracefulShutdown) WaitForShutdown(ctx context.Context) {
	waitCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Waiting for termination signal
	<-waitCtx.Done()

	fmt.Println("Received SIGTERM signal. Gracefully terminating application")
	gs.shutdown()

	componentsToWait := len(gs.components)
	startTime := time.Now()

	for true {
		// Waiting for 200ms
		time.Sleep(200 * time.Millisecond)

		for componentName, cSignal := range gs.components {
			if !cSignal.done {
				select {
				case err := <-cSignal.doneChan:
					if err == nil {
						fmt.Println(fmt.Sprintf("Component %s shutdown gracefully", componentName))
					} else {
						fmt.Println(fmt.Sprintf("Component %s did not shutdown gracefully: %s", componentName, err))
					}
					cSignal.done = true
					componentsToWait--
				default:
					// Do nothing. We go to the next component
				}
			}
		}

		if componentsToWait <= 0 {
			fmt.Println("App gracefully shutdown. Bye!")
			return
		}

		elapsed := time.Now().Sub(startTime)
		if elapsed >= MaxWaitDuration {
			fmt.Println("App took too long to shutdown. Done waiting for graceful shutdown")
			return
		}
	}
}
