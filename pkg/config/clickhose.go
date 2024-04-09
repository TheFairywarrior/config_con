package config

import "github.com/thefairywarrior/config_con/pkg/engines"


type ClickhouseConfig struct {
	name string
	engines.ClickhouseEngine
}

func (c *ClickhouseConfig) Load() (any, error) {
	return c.ClickhouseEngine, nil
}

func (c *ClickhouseConfig) Validate() error {
	return nil
}

func (c *ClickhouseConfig) Name() string {
	return c.name
}


func NewClickhouseConfig(data map[string]any) Configuration {
	return &ClickhouseConfig{
		name: data["name"].(string),
		ClickhouseEngine: engines.NewClickhouseEngine(data["host"].(string), data["port"].(int), data["database"].(string), data["username"].(string), data["password"].(string)),
	}
}