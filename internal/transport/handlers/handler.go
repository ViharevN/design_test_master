package handlers

import (
	"github.com/ViharevN/design_test_master/internal/model"
	"github.com/ViharevN/design_test_master/internal/pkg/errors"
	"github.com/ViharevN/design_test_master/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type handler struct {
	orders service.OrderService
	rooms  service.RoomService
}

func (c *handler) CreateRoom(ctx *gin.Context) {
	var newRoom model.Room
	if err := ctx.ShouldBindJSON(&newRoom); err != nil {
		errors.ApiBadRequestError(ctx, err)
		return
	}

	if err := c.rooms.CreateNewRoom(ctx, newRoom); err != nil {
		errors.ApiInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, newRoom)
}

func (c *handler) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		errors.ApiBadRequestError(ctx, err)
		return
	}

	if err := c.orders.CreateOrder(ctx, newOrder); err != nil {
		errors.ApiInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, newOrder)
}

func (c *handler) GetAvalailableRoomsByDay(ctx *gin.Context) {
	dateParam := ctx.Query("date")
	if dateParam == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Не указана дата"})
		return
	}

	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		errors.ApiBadRequesErrorWithMessage(ctx, "Неверный формат даты")
		return
	}

	rooms, err := c.rooms.GetAvailableRooms(ctx, date)
	if err != nil {
		errors.ApiInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rooms)

}

func (c *handler) GetAvailableRoomsByDateRange(ctx *gin.Context) {
	fromDateParam := ctx.Query("from_date")
	toDateParam := ctx.Query("to_date")
	if fromDateParam == "" || toDateParam == "" {
		errors.ApiBadRequesErrorWithMessage(ctx, "Не указаны даты")
		return
	}

	fromDate, err := time.Parse("2006-01-02", fromDateParam)
	if err != nil {
		errors.ApiBadRequesErrorWithMessage(ctx, "Неверный формат начальной даты")
		return
	}
	toDate, err := time.Parse("2006-01-02", toDateParam)
	if err != nil {
		errors.ApiBadRequesErrorWithMessage(ctx, "Неверный формат конечной даты")
		return
	}

	rooms, err := c.rooms.GetAvailableRoomsByDateRange(ctx, fromDate, toDate)
	if err != nil {
		errors.ApiInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rooms)
}

func NewOrderController(orders service.OrderService, rooms service.RoomService) *handler {
	return &handler{
		orders: orders,
		rooms:  rooms,
	}
}
