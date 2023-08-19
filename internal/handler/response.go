package handler

import (
	"github.com/Vasiliy82/PoolingRouterEmul/internal/logger"
	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Logger().Error(message)
	c.AbortWithStatusJSON(statusCode, error{Message: message})
}
