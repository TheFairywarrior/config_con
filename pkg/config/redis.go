package config

type RedisConfig struct {
	name     string
	host     string
	port     int
	database int
	channel  string
	password string
}

func (c *RedisConfig) Load() (any, error) {
	return nil, nil
}

func (c *RedisConfig) Validate() error {
	return nil
}

func (c *RedisConfig) Name() string {
	return c.name
}

func NewRedisConfig(data map[string]any) Configuration {
	return &RedisConfig{
		name:     data["name"].(string),
		host:     data["host"].(string),
		port:     data["port"].(int),
		database: data["database"].(int),
		channel:  data["channel"].(string),
		password: data["password"].(string),
	}
}
