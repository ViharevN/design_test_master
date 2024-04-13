package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ViharevN/design_test_master/internal/transport"
)

func NewRouter(handler *gin.Engine, order transport.OrderHandler) *gin.Engine {
	router := gin.Default()
	router.Group("/booking")
	{
		handler.POST("/order/create", order.CreateOrder)
		handler.POST("/room/create", order.CreateRoom)
		handler.GET("/available/day", order.GetAvalailableRoomsByDay)
		handler.GET("/available/interval", order.GetAvailableRoomsByDateRange)
	}
	return router
}
