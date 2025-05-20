package main

import (
	"context"
	"net/http"
	"trainer/cmd"
	"trainer/internal/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx := context.Background()
	shutdown := service.NewTracerService(ctx)
	defer shutdown()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()

	grpcServer := cmd.NewGRPCServer(":50051")
	if err := grpcServer.Start(); err != nil {
		panic(err)
	}
}
