package config

import (
	"config_con/pkg/pipe"
	"config_con/pkg/config/consumer"
	"config_con/pkg/config/publisher"
	"config_con/pkg/pipe/transformer"
	"config_con/pkg/utils/environment"
	"context"
	"fmt"
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

// CreatePipelines builds the pipelines from the configuration.
func (config YamlConfiguration) CreatePipelines(cxt context.Context) (map[string]pipe.Pipe, error) {
	consumers := config.Consumers.GetConsumerMap()
	transformers, err := config.Transformers.GetTransformerMap()

	if err != nil {
		return nil, err
	}

	publishers := config.Publishers.GetPublisherMap()

	pipes := make(map[string]pipe.Pipe, len(config.Pipelines))
	for _, pipeline := range config.Pipelines {
		consumer, ok := consumers[pipeline.Consumer]
		if !ok {
			return nil, fmt.Errorf("consumer %s not found", pipeline.Consumer)
		}

		transformer, ok := transformers[pipeline.Transformer]
		if !ok {
			return nil, fmt.Errorf("transformer %s not found", pipeline.Transformer)
		}

		publisher, ok := publishers[pipeline.Publisher]
		if !ok {
			return nil, fmt.Errorf("publisher %s not found", pipeline.Publisher)
		}

		pipe := pipe.NewPipe(
			cxt,
			consumer,
			transformer,
			publisher,
		)
		pipes[pipeline.Name] = pipe
	}
	return pipes, nil
}
