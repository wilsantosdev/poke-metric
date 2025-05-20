package domain

import "testing"

func TestPokemon_ID_Name_Types(t *testing.T) {
	p := NewPokemon(1, "Bulbasaur", []PokemonType{"Grass", "Poison"})
	if p.ID() != 1 {
		t.Errorf("expected ID 1, got %d", p.ID())
	}
	if p.Name() != "Bulbasaur" {
		t.Errorf("expected name Bulbasaur, got %s", p.Name())
	}
	types := p.PokemonTypes()
	if len(types) != 2 || types[0] != "Grass" || types[1] != "Poison" {
		t.Errorf("unexpected types: %v", types)
	}
}
