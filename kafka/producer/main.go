package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	broker := "localhost:9092"
	topic := "test"

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	defer producer.Close()

	// Produce a message to the 'test' topic
	message := "Hello, Kafka from Go!"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	// Wait for message deliveries
	producer.Flush(15 * 1000)
}
