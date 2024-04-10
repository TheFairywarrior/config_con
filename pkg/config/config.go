package config

import "github.com/thefairywarrior/config_con/pkg/pipeline/consumer"

type Configuration interface {
	Load() (any, error)
	Validate() error
	Name() string
}

type ConfigurationLoader interface {
	GetConfig() (map[string]any, error)
}

type Config struct {
	consumerConfigs    map[string]Configuration
	transformerConfigs map[string]Configuration
	publisherConfigs   map[string]Configuration
}

func (c Config) GetConsumer(name string) (consumer.Consumer, error ) {
	con, err := c.consumerConfigs[name].Load()
	if err != nil {
		return nil, err
	}
	return con.(consumer.Consumer), nil
}

func (c Config) GetTransformer(name string) (any, error) {
	return nil, nil
}

func (c Config) GetPublisher(name string) (any, error) {
	return nil, nil
}


var configConstructors = map[string]func(map[string]any) Configuration{
	"redis": NewRedisConfig,
	"clickhouse": NewClickhouseConfig,
}

var loaderConstructors = map[string]func(map[string]any) ConfigurationLoader{
	"yaml": NewYamlLoader,
}


func Register(name string, constructor func(map[string]any) Configuration) {
	configConstructors[name] = constructor
}

func RegisterLoader(name string, constructor func(map[string]any) ConfigurationLoader) {
	loaderConstructors[name] = constructor
}

func parseConfiguration(config map[string]any) map[string]Configuration {
	configurations := make(map[string]Configuration)
	for name, configs := range config {
		for _, c := range configs.([]map[string]any) {
			if _, ok := configConstructors[name]; !ok {
				// TODO add error handling here
				continue
			}
			configInstance := configConstructors[name](c)
			configurations[configInstance.Name()] = configInstance
		}
	}
	return configurations
}

func Load(configs map[string]any) (Config, error) {
	// TODO Add validation for the configs map
	config := Config{
		consumerConfigs:    parseConfiguration(configs["consumer"].(map[string]any)),
		transformerConfigs: parseConfiguration(configs["transformer"].(map[string]any)),
		publisherConfigs:   parseConfiguration(configs["publisher"].(map[string]any)),
	}
	return config, nil
}

