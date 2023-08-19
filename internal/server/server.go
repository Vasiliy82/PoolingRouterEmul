package server

import (
	"net/http"
	"time"

	"github.com/Vasiliy82/PoolingRouterEmul/internal/config"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/logger"
)

const (
	maxHeaderBytes = 1 << 20 // ограничение 1 мб
	ctxTimeout     = 5
)

type Server struct {
	cfg *config.Server
}

func NewServer(cfg *config.Server) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run(handler http.Handler) error {

	httpServer := &http.Server{
		Addr:           s.cfg.Port,
		Handler:        handler,
		ReadTimeout:    time.Second * s.cfg.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		logger.Logger().Infof("Запуск HTTP сервера на порту %s", s.cfg.Port)
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Logger().Fatalf("Ошибка запуска HTTP сервера: %v", err)
		}
	}()

	return nil

}
