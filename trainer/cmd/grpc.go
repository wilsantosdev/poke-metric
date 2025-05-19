package cmd

import (
	"context"
	"fmt"
	"net"
	"trainer/internal/domain"
	pb "trainer/internal/grpc"
	"trainer/internal/repository"
	"trainer/internal/service"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type grpcServer struct {
	port string
	pb.UnimplementedTrainerServer
	trainerService *service.TrainerService
}

func NewGRPCServer(port string) *grpcServer {
	return &grpcServer{
		port: port,
		trainerService: service.NewtrainerService(
			repository.NewTrainerMongoRepository(),
		),
	}
}

func (gs *grpcServer) Start() error {
	lis, err := net.Listen("tcp", gs.port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	pb.RegisterTrainerServer(s, gs)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (gs *grpcServer) CreateTrainer(ctx context.Context, request *pb.CreateTrainerRequest) (*pb.CreateTrainerResponse, error) {
	shutdown := service.NewTracerService(ctx)
	defer shutdown()

	trainer, err := gs.trainerService.CreateTrainer(ctx, request.Name, domain.PokemonType(request.FavoritePokemonType))
	if err != nil {
		return nil, fmt.Errorf("failed to create trainer: %v", err)
	}
	fmt.Printf("Trainer created: %s, %s\n", request.Name, request.FavoritePokemonType)

	var pokemons []*pb.Pokemon
	for _, pokemon := range trainer.Pokemons() {
		var types []string
		for _, t := range pokemon.PokemonTypes() {
			types = append(types, t.String())
		}
		pokemons = append(pokemons, &pb.Pokemon{
			Id:    pokemon.ID(),
			Name:  pokemon.Name(),
			Types: types,
		})
	}

	return &pb.CreateTrainerResponse{
		Id:                  trainer.ID(),
		Name:                trainer.Name(),
		FavoritePokemonType: trainer.FavotitePokemonType().String(),
		Pokemons:            pokemons,
	}, nil
}

func (gs *grpcServer) GetTrainer(ctx context.Context, request *pb.GetTrainerRequest) (*pb.GetTrainerResponse, error) {
	shutdown := service.NewTracerService(ctx)
	defer shutdown()

	trainer, err := gs.trainerService.GetTrainer(ctx, request.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get trainer: %v", err)
	}
	fmt.Printf("Trainer retrieved: %s\n", trainer.Name())

	var pokemons []*pb.Pokemon
	for _, pokemon := range trainer.Pokemons() {
		var types []string
		for _, t := range pokemon.PokemonTypes() {
			types = append(types, t.String())
		}
		pokemons = append(pokemons, &pb.Pokemon{
			Id:    pokemon.ID(),
			Name:  pokemon.Name(),
			Types: types,
		})

		fmt.Println(pokemon)
	}

	return &pb.GetTrainerResponse{
		Id:                  trainer.ID(),
		Name:                trainer.Name(),
		FavoritePokemonType: trainer.FavotitePokemonType().String(),
		Pokemons:            pokemons,
	}, nil
}

func (gs *grpcServer) AddPokemon(ctx context.Context, request *pb.AddPokemonRequest) (*pb.AddPokemonResponse, error) {
	shutdown := service.NewTracerService(ctx)
	defer shutdown()

	var pokemonTypes []domain.PokemonType

	for _, pokemonType := range request.PokemonTypes {
		if domain.PokemonType(pokemonType).IsValid() {
			pokemonTypes = append(pokemonTypes, domain.PokemonType(pokemonType))
		}
	}
	pokemon := domain.NewPokemon(request.PokemonId, request.PokemonName, pokemonTypes)

	_, err := gs.trainerService.AddPokemon(ctx, request.TrainerId, *pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to add pokemon: %v", err)
	}
	fmt.Printf("Pokemon added: %s\n", request.PokemonName)

	return &pb.AddPokemonResponse{
		TrainerId:    request.TrainerId,
		PokemonId:    request.PokemonId,
		PokemonName:  request.PokemonName,
		PokemonTypes: request.PokemonTypes,
	}, nil
}
