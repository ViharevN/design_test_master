package app

import (
	"context"
	"github.com/ViharevN/design_test_master/config"
	repo "github.com/ViharevN/design_test_master/internal/repository/booking"
	"github.com/ViharevN/design_test_master/internal/service"
	"github.com/ViharevN/design_test_master/internal/service/order"
	"github.com/ViharevN/design_test_master/internal/service/room"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type app struct {
	configuration config.Config
	orderService  service.OrderService
	roomService   service.RoomService
	log           *logrus.Logger
}

func NewApp() *app {
	/* **************************** init configuration *************************** */
	loger := logrus.Logger{}
	configuration, err := config.LoadConfig()
	/* **************************** init core pg repos *************************** */
	pgClient, err := repo.NewRepository(configuration)
	if err != nil {
		logrus.Print("DB connect failed")
		panic(err)
	}
	err = pgClient.Ping(context.Background())
	if err != nil {
		loger.Errorf("ping DB failed: %s", err)
	}
	loger.Print("ping DB succeed")

	/* ****************************** init useCases ****************************** */
	serviceOrder := order.NewOrderService(&pgClient)
	serviceRoom := room.NewRoomService(&pgClient)

	return &app{
		configuration: configuration,
		orderService:  serviceOrder,
		roomService:   serviceRoom,
		log:           &loger,
	}
}

func (a *app) Run(ctx context.Context) error {
	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		return a.startServer()
	})

	return group.Wait()
}
