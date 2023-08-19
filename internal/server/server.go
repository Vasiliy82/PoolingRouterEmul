package server

import (
	"net/http"
	"time"

	"github.com/Vasiliy82/PoolingRouterEmul/internal/config"

	"github.com/labstack/echo/v4"
)

const (
	maxHeaderBytes = 1 << 20 // ограничение 1 мб
	ctxTimeout     = 5
)

type Server struct {
	cfg  *config.App
	echo *echo.Echo
}

func NewServer(cfg *config.App) *Server {
	return &Server{
		cfg: cfg,
		echo: echo.New()
	}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:           s.cfg.Port,
		ReadTimeout:    time.Second * s.cfg.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	return nil

}
