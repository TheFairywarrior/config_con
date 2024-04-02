package consumer

type RedisConfig struct {
	Name     string `yaml:"name"`
	Url      string `yaml:"url"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
	Channel  string `yaml:"channel"`
}
