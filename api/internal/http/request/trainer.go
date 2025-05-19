package request

import (
	"encoding/json"
	"fmt"
	"io"
)

type createTrainer struct {
	Name                string `json:"name"`
	FavoritePokemonType string `json:"favorite_pokemon_type"`
}

func NewCreateTrainerRequest(body io.ReadCloser) (*createTrainer, error) {
	var req createTrainer
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (c *createTrainer) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
