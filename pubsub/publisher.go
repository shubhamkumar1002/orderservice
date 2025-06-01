package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"log"
	"orderService/model"
)

func SubmitMessage(paymentEvent model.PaymentEvent) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "primal-device-456513-a8")
	if err != nil {
		log.Fatalf("PubSub Client error: %v", err)
	}
	defer client.Close()

	// Specify the existing topic name.
	topic := client.Topic("orderservice")
	defer topic.Stop()
	messageData, err := json.Marshal(paymentEvent)
	if err != nil {
		log.Fatalf("Failed to marshal order event: %v", err)
	}
	result := topic.Publish(ctx, &pubsub.Message{
		Data: messageData,
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	log.Printf("Message published with ID: %s", id)
}
