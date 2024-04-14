package main

import (
	"context"
	"github.com/ViharevN/design_test_master/internal/app"
)

func main() {
	ctx := context.Background()
	application, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	if err = application.Run(ctx); err != nil {
		panic(err)
	}
}
