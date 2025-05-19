package main

import (
	"api/cmd"
	"context"
)

func main() {

	api := cmd.NewAPI(context.Background(), "8080")
	api.Run()

}
