package service

import (
	"context"
	
	"github.com/ViharevN/design_test_master/internal/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order model.Order) error
}
