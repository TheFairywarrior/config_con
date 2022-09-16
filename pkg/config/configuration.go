package config

import (
	"config_con/pkg/pipe"
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/publisher"
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

// CreatePipelines builds the pipelines from the configuration.
func (config YamlConfiguration) CreatePipelines(cxt context.Context) (map[string]pipe.Pipe, error) {
	consumers := config.Consumers.GetConsumerMap()
	transformers := config.Transformers.GetTransformerMap()
	publishers := config.Publishers.GetPublisherMap()

	pipes := make(map[string]pipe.Pipe, len(config.Pipelines))
	for _, pipeline := range config.Pipelines {
		pipe := pipe.NewPipe(
			cxt,
			consumers[pipeline.Consumer],
			transformers[pipeline.Transformer],
			publishers[pipeline.Publisher],
		)
		pipes[pipeline.Name] = pipe
	}
	return pipes, nil
}
