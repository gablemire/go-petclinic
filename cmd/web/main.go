package main

import (
	"GoPetClinic/src/config"
	"GoPetClinic/src/system"
	"GoPetClinic/src/web"
	"context"
)

func main() {
	appConfig := config.BootstrapConfig()
	gs, appCtx := system.NewGracefulShutdown(context.Background())

	err := system.BootstrapLogger()
	system.PanicOnError(err)

	ginDoneChan, err := web.BootstrapGin(appCtx, appConfig)
	system.PanicOnError(err)
	gs.RegisterComponent("http server", ginDoneChan)

	logger := system.GetLogger("main")
	logger.Info("App ready")

	gs.WaitForShutdown(context.Background())

	system.FlushLogger()
}
