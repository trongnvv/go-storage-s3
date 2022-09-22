package main

import (
	"context"
	"flag"
	"fmt"
	"go-storage-s3/apis/routers"
	"go-storage-s3/common/gracefully"
	"go-storage-s3/common/health"
	"go-storage-s3/common/log"
	"go-storage-s3/common/tracer"
	"go-storage-s3/configs"
	"go-storage-s3/fxloader"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	var pathConfig string
	flag.StringVar(&pathConfig, "config", "configs/config.yaml", "path config")
	flag.Parse()
	configs.LoadConfig(pathConfig)
	log.LoadLogger()
}

func main() {
	var nopLogger fx.Option
	if configs.Get().Mode != "dev" {
		nopLogger = fx.NopLogger
	} else {
		nopLogger = fx.Options()
	}
	app := fx.New(
		fx.Provide(configs.Get),
		fx.Provide(log.GetZeroLog),
		fx.Provide(tracer.InitTracer),
		fx.Provide(health.NewHealthService),
		fx.Provide(health.NewGrpcHealthService),
		fx.Provide(gracefully.NewGracefulShutdownService),
		fx.Options(fxloader.Load()...),
		fx.Invoke(serverLifecycle),
		nopLogger,
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatalf(err, "Err start service")
	}

	// Press Ctrl+C to exit the process
	log.Infof("Press Ctrl+C to exit the process")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-ch

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatalf(err, "Err stop service")
	}
}

func serverLifecycle(lc fx.Lifecycle,
	apiRouter *routers.ApiRouter, cf *configs.Config,
	graceful *gracefully.GracefulShutdownService,
	tp *trace.TracerProvider) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := apiRouter.Engine.Run(fmt.Sprintf("%s", cf.Port)); err != nil {
					log.Fatalf(err, "Cannot start server,address:[%s]", cf.Port)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping restful server.")

			graceful.GracefulStop()
			if err := tp.Shutdown(ctx); err != nil {
				return err
			}
			log.Info("Stopping payment integrator server.")
			return nil
		},
	})
}
