package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		ConnectionTimeout: 30 * time.Second,
		OperationTimeout:  30 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	channel := make(chan pulsar.ConsumerMessage, 100)

	options := pulsar.ConsumerOptions{
		Topic:            "test",
		SubscriptionName: "test-sub",
		Type:             pulsar.Exclusive,
	}

	options.MessageChannel = channel

	consumer, err := client.Subscribe(options)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	for cm := range channel {
		consumer := cm.Consumer
		msg := cm.Message
		fmt.Printf("Consumer %s received a message, msgId: %v, key: %s, content: '%s'\n", consumer.Name(), msg.ID(), string(msg.Key()), string(msg.Payload()))

		consumer.Ack(msg)
	}
}
