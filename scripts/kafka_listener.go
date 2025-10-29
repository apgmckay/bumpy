package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

func main() {
	var broker string
	var topic string
	var groupID string

	var rootCmd = &cobra.Command{
		Use:   "kafka-consumer",
		Short: "A simple Kafka consumer using cobra",
		Run: func(cmd *cobra.Command, args []string) {
			if broker == "" || topic == "" {
				log.Fatalf("Both --broker and --topic flags are required")
			}

			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers: []string{broker},
				Topic:   topic,
				GroupID: groupID, // will be empty if not provided
			})
			defer reader.Close()

			fmt.Printf("Listening on Kafka topic '%s' (broker: %s, groupId: %s)...\n", topic, broker, groupID)

			for {
				msg, err := reader.ReadMessage(context.Background())
				if err != nil {
					log.Fatalf("Error reading message: %v", err)
				}
				fmt.Printf("Received message: %s\n", string(msg.Value))
			}
		},
	}

	rootCmd.Flags().StringVar(&broker, "broker", "", "Kafka broker address (required)")
	rootCmd.Flags().StringVar(&topic, "topic", "", "Kafka topic to consume (required)")
	rootCmd.Flags().StringVar(&groupID, "groupId", "", "Kafka consumer group ID (optional)")

	// mark required flags
	_ = rootCmd.MarkFlagRequired("broker")
	_ = rootCmd.MarkFlagRequired("topic")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
