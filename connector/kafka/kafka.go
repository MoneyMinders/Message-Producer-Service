package kafka

import (
	"Message-Producer-Service/commons"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

const (
	broker = "127.0.0.1:62581"
)

func ProduceMessage(kafkaMessage *commons.KafkaMessage) error {

	topic := "text-messages"
	// Create Kafka producer configuration
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": broker,
	}

	// Create Kafka producer
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
		return err
	}

	// Close the Kafka producer before exiting
	defer producer.Close()

	// Serialize the Kafka message into bytes
	messageBytes, err := json.Marshal(kafkaMessage)
	if err != nil {
		log.Printf("Failed to serialize Kafka message: %v", err)
		return err
	}

	// Create Kafka message object
	responseMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: messageBytes,
	}

	producer.Produce(responseMsg, nil)
	// Checking for errors as Producer produces asynchronously
	for e := range producer.Events() {
		switch event := e.(type) {
		case *kafka.Message:
			if event.TopicPartition.Error != nil {
				log.Fatalf("Failed to produce message: %v", event.TopicPartition.Error)
			}
			log.Printf("Message produced successfully: %s", event.Value)
			return err // Exit the loop after receiving the delivery report
		case kafka.Error:
			log.Fatalf("Error during message production: %v", event)
			return err // Exit the loop in case of an error
		}
	}
	return nil
}
