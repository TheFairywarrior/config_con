package consumer

import (
	"github.com/thefairywarrior/config_con/pkg/api"
	event "github.com/thefairywarrior/config_con/pkg/pipe/consumer"

	"github.com/thefairywarrior/config_con/pkg/pipe/consumer"
)

type ConsumerConfig struct {
	TwitchEventConfigs []TwitchEventConfig `yaml:"twitchConfig"`
	RedisConfigs       []RedisConfig       `yaml:"redisConfig"`
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
	for _, redisConfig := range con.RedisConfigs {
		consumerMap[redisConfig.Name] = consumer.NewRedisConsumer(
			redisConfig.Url,
			redisConfig.Password,
			redisConfig.Database,
			redisConfig.Channel,
		)
	}
	return consumerMap
}
