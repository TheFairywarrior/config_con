package api

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	server        *fiber.App
	currentRoutes map[string]string
}

// serverInstance used to store the server instance in a singleton.
var serverInstance Server

// AddRoute adds the route and the handler to the server.
func (s Server) AddRoute(method, path string, function func(*fiber.Ctx) error) error {

	if _, ok := s.currentRoutes[path]; ok {
		return fmt.Errorf("route %s already exists", path)
	}
	switch method {
	case "GET":
		s.server.Get(path, function)
		return nil
	case "POST":
		s.server.Post(path, function)
		return nil
	}
	return fmt.Errorf("method %s not supported", method)
}

// Run starts the server.
func (s Server) Consume() {
	log.Fatal(s.server.Listen(":3000"))
}

// GetServer returns the server instance.
func GetServer() Server {
	if reflect.DeepEqual(serverInstance, Server{}) {
		serverInstance = Server{server: fiber.New(), currentRoutes: make(map[string]string)}
	}

	return serverInstance
}
