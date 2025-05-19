package domain

import (
	"context"
)

type Trainer struct {
	id                  string
	name                string
	favoritePokemonType PokemonType
	pokemons            []Pokemon
}

func (t Trainer) ID() string {
	return t.id
}
func (t Trainer) Name() string {
	return t.name
}

func (t Trainer) FavotitePokemonType() PokemonType {
	return t.favoritePokemonType
}

func (t Trainer) Pokemons() []Pokemon {
	return t.pokemons
}

func (t *Trainer) AddPokemon(pokemon Pokemon) {
	t.pokemons = append(t.pokemons, pokemon)
}

func NewTrainer(id, name string, favoritePokemonType PokemonType, pokemons []Pokemon) *Trainer {
	return &Trainer{
		id:                  id,
		name:                name,
		favoritePokemonType: favoritePokemonType,
		pokemons:            pokemons,
	}
}

type TrainerRepository interface {
	CreateTrainer(ctx context.Context, name string, favoritePokemonType PokemonType) (*Trainer, error)
	GetTrainer(ctx context.Context, id string) (*Trainer, error)
	AddPokemon(ctx context.Context, id string, pokemon Pokemon) (*Trainer, error)
}
