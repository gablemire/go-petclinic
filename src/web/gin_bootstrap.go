package web

import (
	"GoPetClinic/src/config"
	"GoPetClinic/src/system"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

func BootstrapGin(ctx context.Context, config *config.AppConfig) (system.ShutdownChannel, error) {
	logger := system.GetLogger("gin")
	router := gin.New()

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "healthy")
	})

	router.GET("/ready", func(c *gin.Context) {
		c.String(200, "ready!")
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.HttpServerPort),
		Handler: router,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.Info("HTTP server closed")
			} else {
				logger.Fatal("Could not start the http server", zap.Error(err))
			}
		}
	}()

	logger.Info(fmt.Sprintf("Listening to HTTP traffic at: %s", server.Addr))
	doneChan := make(chan error)

	go func() {
		// Listening for shutdown
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			doneChan <- err
		} else {
			doneChan <- nil
		}
	}()

	return doneChan, nil
}
