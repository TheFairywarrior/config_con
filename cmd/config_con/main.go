package main

import (
	"context"
	"log"

	"github.com/thefairywarrior/config_con/pkg/api"
	"github.com/thefairywarrior/config_con/pkg/config"

	"golang.org/x/exp/maps"
)

func main() {
	log.Println("Starting github.com/thefairywarrior/config_con")
	yamlConfig, err := config.ReadConfiguration()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	pipes, err := yamlConfig.CreatePipelines(ctx)
	if err != nil {
		cancel()
		panic(err)
	}

	server := api.GetServer()
	for _, pipe := range maps.Values(pipes) {
		pipe.Start()
	}

	api.WhenReady()

	log.Println("Starting server")
	server.StartServer(ctx)
}
