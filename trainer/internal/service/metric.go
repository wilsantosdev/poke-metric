package service

import "github.com/prometheus/client_golang/prometheus"

var (
	TrainersCreated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "trainers_created_total",
			Help: "Total de treinadores criados.",
		},
	)

	PokemonsAdded = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pokemons_added_total",
			Help: "Total de pokémons adicionados por trainer e pokemon.",
		},
		[]string{"trainer_id", "pokemon_id"},
	)
)

func init() {
	prometheus.MustRegister(
		TrainersCreated,
		PokemonsAdded,
	)
}
