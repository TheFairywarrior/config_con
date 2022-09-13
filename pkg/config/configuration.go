package config

import (
	"config_con/pkg/pipe"
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/transformer"
	"config_con/pkg/utils/environment"
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
	Transformers []transformer.TransformerConfig `yaml:"transformers"`
	Pipelines    []pipe.PipeConfig               `yaml:"pipelines"`
}
