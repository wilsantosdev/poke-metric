package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hunt/internal/domain"
	"math/rand"
	"net/http"
)

const (
	POKEMON_URL = "https://pokeapi.co/api/v2/pokemon/%v"
)

type PokemonResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type pokemonService struct {
}

func NewPokemonService() *pokemonService {
	return &pokemonService{}
}

func (p *pokemonService) GetPokemonByID(ctx context.Context, pokemonId int32) (*domain.Pokemon, error) {
	httpClient := &http.Client{}
	url := fmt.Sprintf(POKEMON_URL, pokemonId)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return nil, err
	}

	var pokemonResponse PokemonResponse
	err = json.NewDecoder(resp.Body).Decode(&pokemonResponse)
	if err != nil {
		return nil, err
	}
	pokemonTypes := make([]domain.PokemonType, len(pokemonResponse.Types))
	for i, t := range pokemonResponse.Types {
		pokemonTypes[i] = domain.PokemonType(t.Type.Name)
	}
	if !pokemonTypes[0].IsValid() {
		return nil, fmt.Errorf("invalid pokemon type: %s", pokemonTypes[0])
	}
	pokemon := domain.NewPokemon(int32(pokemonResponse.ID), pokemonResponse.Name, pokemonTypes)
	return pokemon, nil
}

func (p *pokemonService) Hunt(ctx context.Context, trainerId, favorite_pokemon_type string) (*domain.Pokemon, error) {
	pokemon, err := p.GetPokemonByID(ctx, p.randomPokemon())
	if err != nil {
		return nil, err
	}
	fmt.Printf("Hunting for Pokemon %s with ID %d\n", pokemon.Name(), pokemon.ID())
	return pokemon, nil
}

func (p *pokemonService) randomPokemon() int32 {
	return int32(rand.Intn(151) + 1)
}
