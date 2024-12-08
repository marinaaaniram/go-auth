package main

import (
	"context"
	"log"
	"time"

	"github.com/marinaaaniram/go-auth/internal/app"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to init app: %s", err.Error())
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalf("Failed to run app: %s", err.Error())
	}
}
