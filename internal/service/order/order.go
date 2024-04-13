package order

import (
	"context"
	"fmt"
	"github.com/ViharevN/design_test_master/internal/repository"

	"github.com/ViharevN/design_test_master/internal/model"
)

type order struct {
	repository repository.Repo
}

func (s *order) CreateOrder(ctx context.Context, order model.Order) error {
	tx, err := s.repository.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to bigin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	available, err := s.repository.CheckRoomAvailability(ctx, tx, order)
	if err != nil {
		return fmt.Errorf("failed check room availability: %s", err)
	}
	if !available {
		return fmt.Errorf("room is not available: %w", err)
	}

	if err = s.repository.CreateOrder(ctx, tx, order); err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func NewOrderService(repository repository.Repo) *order {
	return &order{
		repository: repository,
	}
}
