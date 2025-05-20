package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hunt/internal/domain"
	"math/rand"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	HUNT_TOPIC = "hunt"
)

type HuntMessage struct {
	TrainerID           string `json:"trainer_id"`
	FavoritePokemonType string `json:"favorite_pokemon_type"`
	Attempts            int32  `json:"attempts"`
	MaxAttempts         int32  `json:"max_attempts"`
}

type huntService struct {
	pokemonService domain.PokemonService
	trainerService domain.TrainerService
	kafkaProducer  *kafka.Producer
}

func NewHuntService() *huntService {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"group.id":          "hunt",
	})
	if err != nil {
		panic(err)
	}
	return &huntService{
		pokemonService: NewPokemonService(),
		trainerService: NewTrainerService(),
		kafkaProducer:  kafkaProducer,
	}
}

func (h *huntService) HuntPokemon(ctx context.Context, trainerId string, favorite_pokemon_type string, attemps int32, maxAttempts int32) error {

	if attemps >= maxAttempts {
		return nil
	}

	pokemon, err := h.pokemonService.Hunt(ctx, trainerId, favorite_pokemon_type)
	if err != nil {
		return err
	}

	catched := h.wasCaptured(pokemon, favorite_pokemon_type)

	if catched {
		fmt.Printf("Trainer %s captured a %s!\n", trainerId, pokemon.Name())
		err = h.trainerService.AddPokemon(ctx, trainerId, *pokemon)
		if err != nil {
			fmt.Printf("Error adding pokemon to trainer %s: %v\n", trainerId, err)
		}

	} else {
		fmt.Printf("Trainer %s failed to capture a %s.\n", trainerId, pokemon.Name())
	}

	HuntTotal.WithLabelValues(trainerId, fmt.Sprint(pokemon.ID()), fmt.Sprintf("%t", catched), fmt.Sprintf("%d", attemps)).Inc()

	if !h.produceNewHuntMessage(trainerId, favorite_pokemon_type, attemps, maxAttempts) {
		fmt.Printf("Error producing new hunt message for trainer %s\n", trainerId)
		return fmt.Errorf("error producing new hunt message for trainer %s", trainerId)
	}

	return nil
}

func (h *huntService) wasCaptured(pokemon *domain.Pokemon, favoritePokemonType string) bool {
	isFavoriteType := false
	for _, t := range pokemon.PokemonTypes() {
		if t.String() == favoritePokemonType {
			isFavoriteType = true
			break
		}
	}

	if isFavoriteType {
		return true
	}

	return rand.Intn(100) < 50
}

func (h *huntService) produceNewHuntMessage(trainerId string, favoritePokemonType string, attempts int32, maxAttempts int32) bool {
	huntTopic := HUNT_TOPIC
	msg := HuntMessage{
		TrainerID:           trainerId,
		FavoritePokemonType: favoritePokemonType,
		Attempts:            attempts + 1,
		MaxAttempts:         maxAttempts,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("Error marshalling hunt message: %v\n", err)
		return false
	}
	h.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &huntTopic, Partition: kafka.PartitionAny},
		Value:          jsonData,
	}, nil)

	return true
}
