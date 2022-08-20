package api

import "config_con/pkg/pipe/consumer/twitch"


type ApiConfiguration struct {
	twitchConsumers []twitch.TwitchEventConfig
}
