package twitch

import (
	"config_con/pkg/pipe/queue"

	"github.com/gofiber/fiber/v2"
)

type TwitchEventConfig struct {
	Name        string `yaml:"name"`
	EventSecret string `yaml:"eventSecret"`
	Url         string `yaml: "url"`
}

type TwitchEventConsumer struct {
	TwitchEventConfig
	queue queue.TransformerQueue
}

func (con TwitchEventConfig) Route(cxt *fiber.Ctx) error {
	return cxt.SendString("Hello World")
}
