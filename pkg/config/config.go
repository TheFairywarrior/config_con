package config

type Configuration interface {
	Load() (any, error)
	Validate() error
	Name() string
	Type() string // Can be consumer, step, publisher or pipeline
}

type ConfigurationLoader interface {
	GetConfig() (map[string]any, error)
}

type Config struct {
	consumerConfigs    map[string]Configuration
	transformerConfigs map[string]Configuration
	publisherConfigs   map[string]Configuration
}

var configInstances = map[string]Configuration{}
var configConstructors = map[string]func(map[string]any) Configuration{}
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
