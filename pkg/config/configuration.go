package config

import (
	"config_con/pkg/pipe"
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/publisher"
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer"
	"config_con/pkg/utils/environment"
	"context"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadConfiguration() (YamlConfiguration, error) {
	configFilePath := environment.Settings.ConfigFilePath
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	var configMap YamlConfiguration
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return YamlConfiguration{}, err
	}
	err = yaml.Unmarshal(yamlFile, &configMap)
	if err != nil {
		return YamlConfiguration{}, err
	}
	return configMap, nil
}

type YamlConfiguration struct {
	Consumers    consumer.ConsumerConfig       `yaml:"consumers"`
	Transformers transformer.TransformerConfig `yaml:"transformers"`
	Publishers   publisher.PublisherConfig     `yaml:"publishers"`
	Pipelines    []pipe.PipeConfig             `yaml:"pipelines"`
}

func (config YamlConfiguration) CreatePipelines(cxt context.Context) error {
	consumers := config.Consumers.GetConsumerMap()
	transformers := config.Transformers.GetTransformerMap()
	publishers := config.Publishers.GetPublisherMap()

	for _, pipeline := range config.Pipelines {
		err := pipe.NewPipe(
			cxt,
			pipeline.Name,
			queue.LocalQueue{},
			consumers[pipeline.Consumer],
			transformers[pipeline.Transformer],
			queue.LocalQueue{},
			publishers[pipeline.Publisher],
		)
		if err != nil {
			return err
		}
	}
	return nil
}
