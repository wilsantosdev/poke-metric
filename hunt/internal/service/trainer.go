package service

import (
	"context"
	"fmt"

	"hunt/internal/domain"
	pb "hunt/internal/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	TrainerServiceName = "trainer"
	TrainerServicePort = "50051"
)

type trainerService struct {
	grpcClient pb.TrainerClient
}

func NewTrainerService() *trainerService {
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", TrainerServiceName, TrainerServicePort),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to trainer service: %v", err))
	}
	client := pb.NewTrainerClient(conn)

	return &trainerService{
		grpcClient: client,
	}
}

func (s *trainerService) AddPokemon(ctx context.Context, trainerId string, pokemon domain.Pokemon) error {
	req := &pb.AddPokemonRequest{
		TrainerId:   trainerId,
		PokemonId:   pokemon.ID(),
		PokemonName: pokemon.Name(),
		PokemonTypes: func() []string {
			types := make([]string, len(pokemon.PokemonTypes()))
			for i, t := range pokemon.PokemonTypes() {
				types[i] = string(t)
			}
			return types
		}(),
	}

	_, err := s.grpcClient.AddPokemon(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to add pokemon: %v", err)
	}

	return nil
}
