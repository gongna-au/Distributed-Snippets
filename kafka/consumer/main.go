package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	broker := "localhost:9092"
	groupId := "myGroup"
	topic := "test"

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupId,
		"auto.offset.reset": "earliest", // Use "latest" to only receive new messages
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	defer consumer.Close()

	// Subscribe to the topic
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topic: %s\n", err)
		os.Exit(1)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}
