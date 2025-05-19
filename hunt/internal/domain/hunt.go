package domain

import "context"

type HuntService interface {
	HuntPokemon(ctx context.Context, trainerId string, favorite_pokemon_type string, attemps int32, maxAttempts int32) error
}
