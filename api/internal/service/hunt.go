package service

import (
	"api/internal/domain"
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	huntTopic   = "hunt"
	MaxAttempts = 5
)

type huntService struct {
	kafkaProducer *kafka.Producer
}

type huntMessage struct {
	TrainerID           string `json:"trainer_id"`
	FavoritePokemonType string `json:"favorite_pokemon_type"`
	Atttempts           int    `json:"attempts"`
	MaxAttempts         int    `json:"max_attempts"`
}

func NewHuntService() *huntService {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"group.id":          "api",
	})
	if err != nil {
		panic(err)
	}
	return &huntService{
		kafkaProducer: kafkaProducer,
	}
}

func (hs *huntService) HuntPokemon(ctx context.Context, trainer domain.Trainer) error {
	var huntTopicPtr = func() *string { s := huntTopic; return &s }()

	jsonData, err := json.Marshal(huntMessage{
		TrainerID:           trainer.ID(),
		FavoritePokemonType: trainer.FavotitePokemonType(),
		Atttempts:           0,
		MaxAttempts:         MaxAttempts,
	})
	if err != nil {
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: huntTopicPtr, Partition: kafka.PartitionAny},
		Value:          jsonData,
	}

	err = hs.kafkaProducer.Produce(msg, nil)
	if err != nil {
		return err
	}

	return nil
}
