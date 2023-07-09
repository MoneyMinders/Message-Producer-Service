package controller

import (
	"context"
	"log"
	"net/http"

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

	// Start the server with the router
	log.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}
	return nil
}
