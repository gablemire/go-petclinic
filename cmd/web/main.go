package main

import (
	"GoPetClinic/src/system"
	"context"
)

func main() {
	gs, appCtx := system.NewGracefulShutdown(context.Background())

	logDoneChan, err := system.BootstrapLogger(appCtx)
	system.PanicOnError(err)
	gs.RegisterComponent("logger", logDoneChan)

	logger := system.GetLogger("main")
	logger.Info("App ready")

	gs.WaitForShutdown(context.Background())
}
