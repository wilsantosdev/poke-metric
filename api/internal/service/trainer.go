package service

import (
	"context"
	"fmt"

	"api/internal/domain"
	pb "api/internal/grpc"

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
func (s *trainerService) CreateTrainer(ctx context.Context, name, favoritePokemonType string) (*domain.Trainer, error) {

	req := &pb.CreateTrainerRequest{
		Name:                name,
		FavoritePokemonType: favoritePokemonType,
	}

	resp, err := s.grpcClient.CreateTrainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create trainer: %v", err)
	}

	return domain.NewTrainer(resp.Id, resp.Name, domain.PokemonType(resp.FavoritePokemonType), []domain.Pokemon{}), nil
}

func (s *trainerService) GetTrainer(ctx context.Context, id string) (*domain.Trainer, error) {
	req := &pb.GetTrainerRequest{
		Id: id,
	}

	resp, err := s.grpcClient.GetTrainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get trainer: %v", err)
	}

	pokemons := make([]domain.Pokemon, len(resp.Pokemons))
	for i, p := range resp.Pokemons {
		types := make([]domain.PokemonType, len(p.Types))
		for j, t := range p.Types {
			types[j] = domain.PokemonType(t)
		}
		pokemons[i] = *domain.NewPokemon(p.Id, p.Name, types)
	}

	return domain.NewTrainer(resp.Id, resp.Name, domain.PokemonType(resp.FavoritePokemonType), pokemons), nil
}
