package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TheDigitalMadness/notifications-service-go/internal/config"
	httpController "github.com/TheDigitalMadness/notifications-service-go/internal/controller/http"
	"github.com/TheDigitalMadness/notifications-service-go/internal/repository"
	"github.com/TheDigitalMadness/notifications-service-go/internal/service"
	"github.com/TheDigitalMadness/notifications-service-go/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func main() {
	module := fx.Options(
		fx.Provide(config.MustMakeConfig),
		fx.Provide(httpController.New),
		fx.Provide(service.New),
		fx.Provide(httpController.NewRouter),
		fx.Provide(db.NewPostgres),
		fx.Provide(repository.New),
		fx.Invoke(initLifecycle),
	)

	fx.New(module).Run()
}

func initLifecycle(cfg *config.Config, lc fx.Lifecycle, router *gin.Engine, pool *pgxpool.Pool) {
	addr := fmt.Sprintf(":%d", cfg.HTTP.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	onStart := func(ctx context.Context) error {
		errChan := make(chan error, 1)

		startServer := func(ch chan error) {
			if err := server.ListenAndServe(); err != nil &&
				err != http.ErrServerClosed {
				ch <- err
			}
		}

		go startServer(errChan)

		select {
		case err := <-errChan:
			return err
		case <-time.After(100 * time.Millisecond):
			return nil
		}
	}
	onStop := func(ctx context.Context) error {
		pool.Close()
		return server.Shutdown(ctx)
	}

	lc.Append(
		fx.Hook{
			OnStart: onStart,
			OnStop:  onStop,
		},
	)
}
