package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satanaroom/L0/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Метод инициализации роутинга
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	orders := router.Group("/orders")
	{
		// Метод GET загружает HTML-страницу
		orders.GET("/", h.CreateHTML)
		// Метод POST посылает id заказа для поиска его в кэше
		orders.POST("/", h.GetModel)
	}

	return router
}
