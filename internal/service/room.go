package service

import (
	"context"
	"github.com/ViharevN/design_test_master/internal/model"
	"time"
)

type RoomService interface {
	CreateNewRoom(ctx context.Context, room model.Room) error
	GetAvailableRooms(ctx context.Context, date time.Time) (model.Availability, error)
	GetAvailableRoomsByDateRange(ctx context.Context, from, to time.Time) (model.Availability, error)
}
