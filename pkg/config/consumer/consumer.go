package consumer

import (
	"config_con/pkg/api"
	event "config_con/pkg/pipe/consumer/twitch"

	"config_con/pkg/pipe/consumer"
)

type ConsumerConfig struct {
	TwitchEventConfigs []TwitchEventConfig `yaml:"twitchConfig"`
}

func (con ConsumerConfig) GetConsumerMap() map[string]consumer.Consumer {
	consumerMap := make(map[string]consumer.Consumer)
	api.InitRoutes(len(con.TwitchEventConfigs))
	for _, twitchEventConfig := range con.TwitchEventConfigs {
		consumerMap[twitchEventConfig.Name] = event.NewTwitchEventConsumer(
			twitchEventConfig.Name,
			twitchEventConfig.EventSecret, twitchEventConfig.Url,
		)
	}
	return consumerMap
}
