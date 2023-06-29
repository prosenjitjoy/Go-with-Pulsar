package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "test",
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

		time.Sleep(5 * time.Second)
	}
}
