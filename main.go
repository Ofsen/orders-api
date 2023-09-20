package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Ofsen/orders-api/app"
)

func main() {
	app := app.New(app.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

	cancel()
}