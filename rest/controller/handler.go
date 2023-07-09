package controller

import (
	"Message-Producer-Service/commons"
	"Message-Producer-Service/connector/kafka"
	"Message-Producer-Service/rest/request"
	"encoding/json"
	"log"
	"net/http"
)



func handleProcessMessageRequest(w http.ResponseWriter, r *http.Request) {
	
		// Get the request body
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
	
		// Decode the request body into struct and check for errors
		var request request.ProcessMessageRequest
		err := decoder.Decode(&request)
		if err != nil {
			log.Printf("Failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		// Validate the request body (if needed)
		// ...
	
		kafkaMessage := &commons.KafkaMessage{
			Category: request.Category,
			AppName:  request.AppName,
			Message:  request.Message,
		}
		
		// Produce the message to Kafka
		err = kafka.ProduceMessage(kafkaMessage)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		// Send the response Request Accepted
		w.WriteHeader(http.StatusAccepted)
	}
