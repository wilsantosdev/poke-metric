package main

import (
	"hunt/cmd"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8082", nil)
	}()
	cmd.NewConsumer()
}
