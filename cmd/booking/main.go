package main

import (
	"context"
	"github.com/ViharevN/design_test_master/internal/app"
)

func main() {
	ctx := context.Background()
	application := app.NewApp()
	if err := application.Run(ctx); err != nil {
		panic(err)
	}
}
