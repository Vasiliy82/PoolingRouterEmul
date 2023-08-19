package service

import (
	"github.com/Vasiliy82/PoolingRouterEmul/internal/model"
)

type PoolingRouterService interface {
	Request(*model.Request) (*model.Response, error)
	Results(*model.Response) (*model.Results, error)
}

type Service struct {
	PoolingRouterService
}

func NewService(poolingRouterService PoolingRouterService) *Service {
	return &Service{PoolingRouterService: poolingRouterService}
}
