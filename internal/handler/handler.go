package handler

import (
	"github.com/Vasiliy82/PoolingRouterEmul/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		request := api.Group("/pooling")
		{
			request.POST("/request", h.request)
		}

	}
	return router
}
