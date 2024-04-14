package app

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ViharevN/design_test_master/internal/transport/handlers"
	"github.com/ViharevN/design_test_master/internal/transport/router"
)

func (a *app) startServer() error {
	gin.SetMode(a.configuration.Debug.GinDebugMode)

	handler := gin.Default()

	/* ************************* init http controllers *************************** */
	controller := handlers.NewOrderController(a.orderService, a.roomService)
	/* *************************************************************************** */

	/* ******************** register controllers in router *********************** */
	router.NewRouter(handler, controller)
	/* *************************************************************************** */

	listener := &http.Server{
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	server, err := net.Listen("tcp", fmt.Sprintf(":%s", a.configuration.ListenerHttpPort))
	if err != nil {
		panic(err)
	}
	a.log.Print("HTTP Server is listening port: ", a.configuration.ListenerHttpPort)
	return listener.Serve(server)
}
