package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"hunt/internal/service"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	huntTopic = "hunt"
)

func NewConsumer() {

	huntService := service.NewHuntService()

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"auto.offset.reset": "earliest",
		"group.id":          "hunt-consumer",
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{huntTopic}, nil)

	if err != nil {
		panic(err)
	}
	var run = true

	for run {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			var huntMessage service.HuntMessage
			err := json.Unmarshal(e.Value, &huntMessage)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", err)
				run = false
				break
			}

			err = huntService.HuntPokemon(context.Background(), huntMessage.TrainerID, huntMessage.FavoritePokemonType, huntMessage.Attempts, huntMessage.MaxAttempts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", err)
				run = false
				break
			}
			fmt.Printf("Consumed message: %s\n", string(e.Value))
			topic := huntTopic
			consumer.CommitOffsets([]kafka.TopicPartition{{
				Topic:     &topic,
				Partition: e.TopicPartition.Partition,
				Offset:    e.TopicPartition.Offset + 1,
			}})

		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	consumer.Close()

}
