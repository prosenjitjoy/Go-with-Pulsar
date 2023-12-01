package main

import (
	"context"
	"fmt"
	"log"
	"main/utils"
	"math/rand"
	"time"

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

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: cfg.TopicName,
	})
	if err != nil {
		log.Fatal(err)
	}

	users := []string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := []string{"book", "alarm clock", "t-shirt", "gift card", "batteries"}

	for {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]

		msgId, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(data),
			Key:     key,
		})
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("Produced event to topic %s: key = %s value = %s msgId = %s\n", "test", string(key), string(data), string(msgId.String()))
		}

		time.Sleep(cfg.Delay)
	}
}
