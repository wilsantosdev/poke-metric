package domain

import (
	"context"
)

type HuntService interface {
	HuntPokemon(ctx context.Context, trainer Trainer, maxAttemps int32) error
}
