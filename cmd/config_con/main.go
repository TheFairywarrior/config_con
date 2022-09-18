package main

import (
	"config_con/pkg/api"
	"config_con/pkg/config"
	"context"
	"log"
	"time"

	"golang.org/x/exp/maps"
)

func main() {
	log.Println("Starting config_con")
	yamlConfig, err := config.ReadConfiguration()
	if err != nil {
		panic(err)
	}
	cxt, cancel := context.WithCancel(context.Background())
	pipes, err := yamlConfig.CreatePipelines(cxt)
	if err != nil {
		cancel()
		panic(err)
	}

	server := api.GetServer()
	for _, pipe := range maps.Values(pipes) {
		pipe.Start()
	}
	// A simple fix to stop the server running before the api paths have been added.
	// TODO: Replace this sleep with a check that all of the consumers are ready.
	time.Sleep(10 * time.Second) 
	server.StartServer(cxt)
}
