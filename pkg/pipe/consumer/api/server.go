package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


type Server struct {
	server *fiber.App
}

// serverInstance used to store the server instance in a singleton.
var serverInstance Server

// AddRoute adds the route and the handler to the server.
func (s Server) AddRoute(method, path string, function func(*fiber.Ctx) error) {
	switch method {
	case "GET":
		s.server.Get(path, function)
		break
	case "POST":
		s.server.Post(path, function)
		break
	}
}

// Run starts the server.
func (s Server) Consume() {
	log.Fatal(s.server.Listen(":3000"))
}

// GetServer returns the server instance.
func GetServer() Server {
	if serverInstance == (Server{}) {
		serverInstance = Server{server: fiber.New()}
	}

	return serverInstance
}
