package handler

import (
	"net/http"

	"github.com/Vasiliy82/PoolingRouterEmul/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) request(c *gin.Context) {
	var request model.Request

	if err := c.BindJSON(&request); err != nil {
		// logger.Logger().Errorf("Ошибка при обработке запроса: ", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.service.PoolingRouterService.Request(&request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, *response)

}
