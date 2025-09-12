package main

import (
	"context"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("script take 1 argument for topic name.")
		os.Exit(1)
	}

	topicName := os.Args[1]

	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer admin.Close()

	topic := kafka.TopicSpecification{
		Topic:             topicName,
		NumPartitions:     3,
		ReplicationFactor: 1,
	}

	results, err := admin.CreateTopics(context.TODO(), []kafka.TopicSpecification{topic})
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Printf("Created topic %s: %v\n", result.Topic, result.Error)
	}
}
