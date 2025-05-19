package domain

import "context"

type Trainer struct {
	id                  string
	name                string
	favoritePokemonType string
	pokemons            []Pokemon
}

func (t Trainer) ID() string {
	return t.id
}
func (t Trainer) Name() string {
	return t.name
}

func (t Trainer) FavotitePokemonType() string {
	return t.favoritePokemonType
}

func (t Trainer) Pokemons() []Pokemon {
	return t.pokemons
}

func (t *Trainer) AddPokemon(pokemon Pokemon) {
	t.pokemons = append(t.pokemons, pokemon)
}

type NewTrainerDTO struct {
	ID                  string
	Name                string
	FavoritePokemonType string
	Pokemons            []Pokemon
}

func NewTrainer(id, name string, favoritePokemonType PokemonType, pokemons []Pokemon) *Trainer {
	return &Trainer{
		id:                  id,
		name:                name,
		favoritePokemonType: string(favoritePokemonType),
		pokemons:            pokemons,
	}
}

type TrainerRepository interface {
	CreateTrainer(ctx context.Context, name, favoritePokemonType string) (*Trainer, error)
}

type TrainerService interface {
	CreateTrainer(ctx context.Context, name, favoritePokemonType string) (*Trainer, error)
	GetTrainer(ctx context.Context, id string) (*Trainer, error)
}
