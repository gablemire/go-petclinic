package system

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

func BootstrapLogger() error {
	if logger != nil {
		return nil
	}

	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		return err
	}

	return nil
}

func GetLogger(component string) *zap.Logger {
	if logger == nil {
		panic("Logger not initialized")
	}

	return logger.With(zap.String("component", component))
}

func FlushLogger() {
	if logger == nil {
		return
	}

	err := logger.Sync()
	if err != nil {
		fmt.Println(fmt.Sprintf("error syncing logger: %v", err))
	}
}
