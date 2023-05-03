package system

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger
var doneChan <-chan error

func BootstrapLogger(ctx context.Context) (<-chan error, error) {
	if logger != nil && doneChan != nil {
		return doneChan, nil
	}

	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		return nil, err
	}

	doneChan := make(chan error)

	go func() {
		<-ctx.Done()

		err := logger.Sync()
		if err == nil {
			doneChan <- nil
		} else {
			doneChan <- fmt.Errorf("error flushing logs: %w", err)
		}
	}()

	return doneChan, nil
}

func GetLogger(component string) *zap.Logger {
	if logger == nil {
		panic("Logger not initialized")
	}

	return logger.With(zap.String("component", component))
}
