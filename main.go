package main

import (
	"Message-Producer-Service/rest/controller"
	"context"
	"log"

	"golang.org/x/sync/errgroup"
)


func main() {
	// Create a new error group
	eg, ctx := errgroup.WithContext(context.Background())

	// Start the server in a goroutine
	eg.Go(func() error {
		return controller.RouterActor(ctx)
	})
	// Wait for any error to occur
	if err := eg.Wait(); err != nil {
		// Handle the error if neede
		log.Fatal("Error in main: ", err)
		panic(err)
	}
}