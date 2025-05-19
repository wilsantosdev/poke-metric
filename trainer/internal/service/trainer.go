package service

import (
	"context"
	"trainer/internal/domain"
)

type TrainerService struct {
	trainerRepository domain.TrainerRepository
}

func NewtrainerService(trainerRepository domain.TrainerRepository) *TrainerService {
	return &TrainerService{
		trainerRepository: trainerRepository,
	}
}
func (t *TrainerService) CreateTrainer(ctx context.Context, name string, favoritePokemonType domain.PokemonType) (*domain.Trainer, error) {
	trainer, err := t.trainerRepository.CreateTrainer(ctx, name, favoritePokemonType)
	if err != nil {
		return nil, err
	}
	return trainer, nil
}

func (t *TrainerService) GetTrainer(ctx context.Context, id string) (*domain.Trainer, error) {
	trainer, err := t.trainerRepository.GetTrainer(ctx, id)
	if err != nil {
		return nil, err
	}
	return trainer, nil
}

func (t *TrainerService) AddPokemon(ctx context.Context, id string, pokemon domain.Pokemon) (*domain.Trainer, error) {
	trainer, err := t.trainerRepository.AddPokemon(ctx, id, pokemon)
	if err != nil {
		return nil, err
	}
	return trainer, nil
}
