package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satanaroom/L0/pkg/service"
)

// "github.com/gin-gonic/gin"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	orders := router.Group("/orders")
	{
		orders.GET("/", h.CreateHTML)
		orders.POST("/", h.GetModel)
	}

	return router
}
