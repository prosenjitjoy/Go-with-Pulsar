package main

import (
	"fmt"
	"log"
	"main/utils"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	cfg := utils.LoadConfig()

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               cfg.PulsarURL,
		ConnectionTimeout: cfg.ConnectionTimeout,
		OperationTimeout:  cfg.OperationTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	channel := make(chan pulsar.ConsumerMessage, 100)

	options := pulsar.ConsumerOptions{
		Topic:            cfg.TopicName,
		SubscriptionName: cfg.SubscriberName,
		Type:             pulsar.Exclusive,
		MessageChannel:   channel,
	}

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
