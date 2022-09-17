package consumer

import (
	"config_con/pkg/pipe/consumer/twitch"
	"config_con/pkg/pipe/queue"
	"context"
)

// Consumer interface is used in the pipeline to consume the data from multiple sources.
type Consumer interface {
	Consume(context.Context, queue.Queue) error
}

type ConsumerConfig struct {
	TwitchEventConfigs []twitch.TwitchEventConsumer `yaml:"twitchConfig"`
}

func (con ConsumerConfig) GetConsumerMap() map[string]Consumer {
	consumerMap := make(map[string]Consumer)
	for _, twitchEventConfig := range con.TwitchEventConfigs {
		consumerMap[twitchEventConfig.Name] = twitchEventConfig
	}
	return consumerMap
}
