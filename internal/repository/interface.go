package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/ViharevN/design_test_master/internal/model"
)

type Repo interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	CreateOrder(ctx context.Context, tx pgx.Tx, order model.Order) error
	CreateRoom(ctx context.Context, room model.Room) error
	GetAvailableRoomsByDate(ctx context.Context, date time.Time) (model.Availability, error)
	GetAvailableRoomsByRange(ctx context.Context, dateFrom, dateTo time.Time) (model.Availability, error)
	CheckRoomAvailability(ctx context.Context, tx pgx.Tx, order model.Order) (bool, error)
}
