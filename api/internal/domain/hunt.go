package domain

import (
	"context"
)

type HuntService interface {
	HuntPokemon(ctx context.Context, trainer Trainer) error
}
