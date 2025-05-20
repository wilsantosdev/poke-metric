package service

import (
	"context"
	"fmt"
	"testing"

	pb "api/internal/grpc"

	"google.golang.org/grpc"
)

type mockTrainerClient struct {
	createTrainerFunc func(ctx context.Context, in *pb.CreateTrainerRequest, opts ...grpc.CallOption) (*pb.CreateTrainerResponse, error)
	getTrainerFunc    func(ctx context.Context, in *pb.GetTrainerRequest, opts ...grpc.CallOption) (*pb.GetTrainerResponse, error)
}

func (m *mockTrainerClient) CreateTrainer(ctx context.Context, in *pb.CreateTrainerRequest, opts ...grpc.CallOption) (*pb.CreateTrainerResponse, error) {
	return m.createTrainerFunc(ctx, in, opts...)
}
func (m *mockTrainerClient) GetTrainer(ctx context.Context, in *pb.GetTrainerRequest, opts ...grpc.CallOption) (*pb.GetTrainerResponse, error) {
	return m.getTrainerFunc(ctx, in, opts...)
}

func (m *mockTrainerClient) AddPokemon(ctx context.Context, in *pb.AddPokemonRequest, opts ...grpc.CallOption) (*pb.AddPokemonResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func TestTrainerService_CreateTrainer_Success(t *testing.T) {
	mockClient := &mockTrainerClient{
		createTrainerFunc: func(ctx context.Context, in *pb.CreateTrainerRequest, opts ...grpc.CallOption) (*pb.CreateTrainerResponse, error) {
			return &pb.CreateTrainerResponse{
				Id:                  "123",
				Name:                in.Name,
				FavoritePokemonType: in.FavoritePokemonType,
			}, nil
		},
	}
	svc := &trainerService{grpcClient: mockClient}
	trainer, err := svc.CreateTrainer(context.Background(), "Ash", "electric")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if trainer.ID() != "123" || trainer.Name() != "Ash" || trainer.FavotitePokemonType() != "electric" {
		t.Errorf("unexpected trainer: %+v", trainer)
	}
}

func TestTrainerService_CreateTrainer_Error(t *testing.T) {
	mockClient := &mockTrainerClient{
		createTrainerFunc: func(ctx context.Context, in *pb.CreateTrainerRequest, opts ...grpc.CallOption) (*pb.CreateTrainerResponse, error) {
			return nil, fmt.Errorf("some error")
		},
	}
	svc := &trainerService{grpcClient: mockClient}
	_, err := svc.CreateTrainer(context.Background(), "Ash", "electric")
	if err == nil || err.Error() != "failed to create trainer: some error" {
		t.Errorf("expected wrapped error, got %v", err)
	}
}

func TestTrainerService_GetTrainer_Success(t *testing.T) {
	mockClient := &mockTrainerClient{
		getTrainerFunc: func(ctx context.Context, in *pb.GetTrainerRequest, opts ...grpc.CallOption) (*pb.GetTrainerResponse, error) {
			return &pb.GetTrainerResponse{
				Id:                  "123",
				Name:                "Ash",
				FavoritePokemonType: "electric",
				Pokemons: []*pb.Pokemon{
					{
						Id:    1,
						Name:  "bulbasaur",
						Types: []string{"grass", "poison"},
					},
				},
			}, nil
		},
	}
	svc := &trainerService{grpcClient: mockClient}
	trainer, err := svc.GetTrainer(context.Background(), "123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if trainer.ID() != "123" || trainer.Name() != "Ash" || trainer.FavotitePokemonType() != "electric" {
		t.Errorf("unexpected trainer: %+v", trainer)
	}
	if len(trainer.Pokemons()) != 1 || trainer.Pokemons()[0].Name() != "bulbasaur" {
		t.Errorf("unexpected pokemons: %+v", trainer.Pokemons())
	}
}

func TestTrainerService_GetTrainer_Error(t *testing.T) {
	mockClient := &mockTrainerClient{
		getTrainerFunc: func(ctx context.Context, in *pb.GetTrainerRequest, opts ...grpc.CallOption) (*pb.GetTrainerResponse, error) {
			return nil, fmt.Errorf("not found")
		},
	}
	svc := &trainerService{grpcClient: mockClient}
	_, err := svc.GetTrainer(context.Background(), "123")
	if err == nil || err.Error() != "failed to get trainer: not found" {
		t.Errorf("expected wrapped error, got %v", err)
	}
}
