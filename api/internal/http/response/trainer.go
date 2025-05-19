package response

import (
	"api/internal/domain"
	"strconv"
)

type pokemonResponse struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Types []string `json:"type"`
}

type createTrainerResponse struct {
	ID                  string            `json:"id"`
	Name                string            `json:"name"`
	FavoritePokemonType string            `json:"favorite_pokemon_type"`
	Pokemons            []pokemonResponse `json:"pokemons"`
}

func NewCreateTrainerResponse(trainer domain.Trainer) createTrainerResponse {
	pokemons := make([]pokemonResponse, len(trainer.Pokemons()))
	for i, p := range trainer.Pokemons() {
		pokemons[i] = pokemonResponse{
			ID:   string(p.ID()),
			Name: p.Name(),
			Types: func() []string {
				types := make([]string, len(p.PokemonTypes()))
				for i, t := range p.PokemonTypes() {
					types[i] = string(t)
				}
				return types
			}(),
		}
	}
	return createTrainerResponse{
		ID:                  trainer.ID(),
		Name:                trainer.Name(),
		FavoritePokemonType: trainer.FavotitePokemonType(),
		Pokemons:            pokemons,
	}
}

func NewGetTrainerResponse(trainer domain.Trainer) createTrainerResponse {
	pokemons := make([]pokemonResponse, len(trainer.Pokemons()))
	for i, p := range trainer.Pokemons() {
		pokemons[i] = pokemonResponse{
			ID:   strconv.Itoa(int(p.ID())),
			Name: p.Name(),
			Types: func() []string {
				types := make([]string, len(p.PokemonTypes()))
				for i, t := range p.PokemonTypes() {
					types[i] = string(t)
				}
				return types
			}(),
		}
	}
	return createTrainerResponse{
		ID:                  trainer.ID(),
		Name:                trainer.Name(),
		FavoritePokemonType: trainer.FavotitePokemonType(),
		Pokemons:            pokemons,
	}
}
