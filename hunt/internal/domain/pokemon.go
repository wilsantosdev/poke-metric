package domain

import (
	"context"
)

type PokemonType string

const (
	Normal   PokemonType = "normal"
	Fighting PokemonType = "fighting"
	Flying   PokemonType = "flying"
	Poison   PokemonType = "poison"
	Ground   PokemonType = "ground"
	Rock     PokemonType = "rock"
	Bug      PokemonType = "bug"
	Ghost    PokemonType = "ghost"
	Steel    PokemonType = "steel"
	Fire     PokemonType = "fire"
	Water    PokemonType = "water"
	Grass    PokemonType = "grass"
	Electric PokemonType = "electric"
	Psychic  PokemonType = "psychic"
	Ice      PokemonType = "ice"
	Dragon   PokemonType = "dragon"
	Dark     PokemonType = "dark"
	Fairy    PokemonType = "fairy"
	Unknown  PokemonType = "unknown"
)

func (p PokemonType) String() string {
	switch p {
	case Normal:
		return "normal"
	case Fighting:
		return "fighting"
	case Flying:
		return "flying"
	case Poison:
		return "poison"
	case Ground:
		return "ground"
	case Rock:
		return "rock"
	case Bug:
		return "bug"
	case Ghost:
		return "ghost"
	case Steel:
		return "steel"
	case Fire:
		return "fire"
	case Water:
		return "water"
	case Grass:
		return "grass"
	case Electric:
		return "electric"
	case Psychic:
		return "psychic"
	case Ice:
		return "ice"
	case Dragon:
		return "dragon"
	case Dark:
		return "dark"
	case Fairy:
		return "fairy"
	default:
		return ""
	}
}
func (p PokemonType) IsValid() bool {
	switch p {
	case Normal, Fighting, Flying, Poison, Ground, Rock, Bug, Ghost, Steel,
		Fire, Water, Grass, Electric, Psychic, Ice, Dragon, Dark, Fairy:
		return true
	default:
		return false
	}
}

type Pokemon struct {
	id           int32
	name         string
	pokemonTypes []PokemonType
}

func NewPokemon(id int32, name string, pokemonTypes []PokemonType) *Pokemon {
	return &Pokemon{
		id:           id,
		name:         name,
		pokemonTypes: pokemonTypes,
	}
}

func (p Pokemon) ID() int32 {
	return p.id
}
func (p Pokemon) Name() string {
	return p.name
}
func (p Pokemon) PokemonTypes() []PokemonType {
	return p.pokemonTypes
}

type PokemonAPI interface {
	GetPokemonByID(id int) (Pokemon, error)
}

type PokemonService interface {
	Hunt(ctx context.Context, trainerId, favorite_pokemon_type string) (*Pokemon, error)
	GetPokemonByID(ctx context.Context, id int32) (*Pokemon, error)
}
