package room

import (
	"context"
	"fmt"
	"github.com/ViharevN/design_test_master/internal/repository"
	"time"

	"github.com/ViharevN/design_test_master/internal/model"
)

type room struct {
	repository repository.Repo
}

func (r room) CreateNewRoom(ctx context.Context, apartment model.Room) error {
	if err := r.repository.CreateRoom(ctx, apartment); err != nil {
		return fmt.Errorf("failed to create new room: %w", err)
	}
	return nil
}

func (r room) GetAvailableRooms(ctx context.Context, date time.Time) (model.Availability, error) {
	rooms, err := r.repository.GetAvailableRoomsByDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get available rooms: %w", err)
	}
	return rooms, nil
}

func (r room) GetAvailableRoomsByDateRange(ctx context.Context, from, to time.Time) (model.Availability, error) {
	rooms, err := r.repository.GetAvailableRoomsByRange(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get available rooms for range: %w", err)
	}
	return rooms, nil
}

func NewRoomService(repository repository.Repo) *room {
	return &room{
		repository: repository,
	}
}
