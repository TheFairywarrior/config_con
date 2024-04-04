package config


type YamlLoader struct {
	file string
}

func (loader YamlLoader) GetConfig() (map[string]any, error) {
	return nil, nil
}

func NewYamlLoader(data map[string]any) ConfigurationLoader {
	return &YamlLoader{
		file: data["file"].(string),
	}
}


