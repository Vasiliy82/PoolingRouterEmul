package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vasiliy82/PoolingRouterEmul/internal/config"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/handler"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/logger"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/repository"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/server"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/service"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/tracer"
	"go.uber.org/zap/zapcore"
)

func main() {

	ctx := context.Background()

	defer func() {
		_ = logger.Logger().Sync()
	}()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadAll()
	if err != nil {
		logger.Logger().Fatalf("ошибка загрузки конфигурации: %v", err)
	}

	logLevel := zapcore.ErrorLevel

	if errParse := logLevel.UnmarshalText([]byte(cfg.Server.LogLevel)); errParse != nil {
		logger.Logger().Errorf("ошибка получения значения LogLevel: %v", errParse)
	}

	logger.SetLevel(logLevel)

	httpTracerShutdown, err := tracer.InitHTTPProvider(cfg.Server.TraceURL, cfg.Server.AppName(), int64(os.Getpid()))
	if err != nil {
		logger.Logger().Fatalf("Ошибка инициализации OpenTrace: %v", err)
	}

	defer func() {
		sCtx := context.Background()
		if err := httpTracerShutdown(sCtx); err != nil {
			log.Fatalf("Ошибка остановки OpenTrace (HTTP TracerProvider): %v", err)
		}
	}()

	service := service.NewService(service.NewPoolingRouterService(repository.NewPoolingStorage()))
	handler := handler.NewHandler(service)
	server := server.NewServer(cfg.Server)
	server.Run(handler.InitRoutes())

	<-ctx.Done()

}
