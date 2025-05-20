package service

import "github.com/prometheus/client_golang/prometheus"

var (
	HuntTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hunt_total",
			Help: "Caçada de pokemons realizada por trainer.",
		},
		[]string{"trainer_id", "pokemon_id", "captured", "attemps"},
	)
)

func init() {
	prometheus.MustRegister(
		HuntTotal,
	)
}
