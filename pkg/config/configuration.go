package config

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/utils/environment"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	consumers map[string]consumer.Consumer
}

type YamlConfiguration struct {
	Consumers []consumer.ConsumerConfig `yaml:"consumers"`
}

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

func ReadConsumerConfigurations() map[string]consumer.Consumer {
	return map[string]consumer.Consumer{}
}
