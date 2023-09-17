package main

import (
	"context"
	"fmt"

	"github.com/Ofsen/orders-api/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}