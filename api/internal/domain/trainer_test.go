package domain

import "testing"

func TestTrainer_Basic(t *testing.T) {

	pokemons := NewPokemon(1, "bullbasaur", []PokemonType{"grass", "poison"})

	trainer := NewTrainer("1", "Ash", "electric", []Pokemon{*pokemons})

	if trainer.ID() != "1" {
		t.Errorf("expected ID 1, got %s", trainer.ID())
	}
	if trainer.Name() != "Ash" {
		t.Errorf("expected name Ash, got %s", trainer.Name())
	}
	if trainer.FavotitePokemonType() != "electric" {
		t.Errorf("expected favorite type Electric, got %s", trainer.FavotitePokemonType())
	}

	if len(trainer.Pokemons()) != 1 {
		t.Errorf("expected 1 pokemon, got %d", len(trainer.Pokemons()))
	}
	if trainer.Pokemons()[0].ID() != 1 {
		t.Errorf("expected pokemon ID 1, got %d", trainer.Pokemons()[0].ID())
	}
	if trainer.Pokemons()[0].Name() != "bullbasaur" {
		t.Errorf("expected pokemon name bullbasaur, got %s", trainer.Pokemons()[0].Name())
	}
}
