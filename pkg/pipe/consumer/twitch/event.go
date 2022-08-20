package twitch

import (
	"config_con/pkg/pipe/queue"

	"github.com/gofiber/fiber/v2"
)

type TwitchEventConfig struct {
	eventSecret string
	route       string
}


type TwitchEventConsumer struct {
	configuration TwitchEventConfig
	queue queue.TransformerQueue
}



func (con TwitchEventConfig) Route(cxt *fiber.Ctx) error {
	return cxt.SendString("Hello World")
}
