package config

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/utils/shortcuts"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	consumers map[string]consumer.Consumer
}


type YamlConfiguration struct {
	consumers []consumer.ConsumerConfig
}

func ReadConfiguration() (shortcuts.Map, error) {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	var configMap shortcuts.Map
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &configMap)
	if err != nil {
		return nil, err
	}
	return configMap, nil
}

func ReadConsumerConfigurations() map[string]consumer.Consumer {
	return map[string]consumer.Consumer{}
}
