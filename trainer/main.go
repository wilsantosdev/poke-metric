package main

import (
	"context"
	"trainer/cmd"
	"trainer/internal/service"
)

func main() {
	ctx := context.Background()
	shutdown := service.NewTracerService(ctx)
	defer shutdown()
	grpcServer := cmd.NewGRPCServer(":50051")
	if err := grpcServer.Start(); err != nil {
		panic(err)
	}
}
