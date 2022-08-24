package api

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

//TestServer is testing the basic building of the server app.
func TestServer(t *testing.T) {
	server := GetServer()
	server.AddRoute("GET", "/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	server.AddRoute("POST", "/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}
