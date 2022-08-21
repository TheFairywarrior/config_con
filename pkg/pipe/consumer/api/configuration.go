package api

import "config_con/pkg/pipe/consumer/twitch"


type ApiConfiguration struct {
	TwitchConsumers []twitch.TwitchEventConfig `yaml:"twitchConsumers"`
}
