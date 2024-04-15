package transport

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	CreateOrder(ctx *gin.Context)
	GetAvalailableRoomsByDay(ctx *gin.Context)
	GetAvailableRoomsByDateRange(ctx *gin.Context)
	CreateRoom(ctx *gin.Context)
}
