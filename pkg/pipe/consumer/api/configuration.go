package api

import (
	"config_con/pkg/pipe/consumer/twitch"

	"github.com/gofiber/fiber/v2"
)


type ApiConfiguration struct {
	TwitchConsumers []twitch.TwitchEventConfig `yaml:"twitchConsumers"`
}

type ApiService interface {
	Route(cxt *fiber.Ctx) error
}

type ApiConsumer struct {
	
}

