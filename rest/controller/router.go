package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	baseRouteProcessMessage = "/rest/v1"
)

func RouterActor(ctx context.Context) error {
	// Create a new router
	router := mux.NewRouter()
	// Define the routes
	router.HandleFunc(baseRouteProcessMessage+"/process-message", handleProcessMessageRequest).Methods("POST")

	// Create a new HTTP server with the router
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		log.Println("Server listening on :8080")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for the cancellation signal from the context
	<-ctx.Done()

	// Shutdown the server gracefully
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("Failed to gracefully shutdown server: %v", err)
		return err
	}

	return nil
}

