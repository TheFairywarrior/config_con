package config

import "github.com/thefairywarrior/config_con/pkg/engines"

type RedisConfig struct {
	name string
	engines.RedisEngine
}

func (c *RedisConfig) Load() (any, error) {
	return c.RedisEngine, nil
}

func (c *RedisConfig) Validate() error {
	return nil
}

func (c *RedisConfig) Name() string {
	return c.name
}

func NewRedisConfig(data map[string]any) Configuration {
	return &RedisConfig{
		name:        data["name"].(string),
		RedisEngine: engines.NewRedisEngine(data["host"].(string), data["port"].(int), data["database"].(int), data["channel"].(string), data["password"].(string)),
	}
}
