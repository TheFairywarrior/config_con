package config

import "github.com/thefairywarrior/config_con/pkg/pipeline/transformer/step"


type MapperStepConfig struct {
	name string
	step.MapperStep
}

func (c *MapperStepConfig) Load() (any, error) {
	return c.MapperStep, nil
}

func (c *MapperStepConfig) Validate() error {
	return nil
}

func (c *MapperStepConfig) Name() string {
	return c.name
}


func NewMapperStepConfig(data map[string]any) Configuration {
	return &MapperStepConfig{
		name:        data["name"].(string),
		MapperStep: step.NewMapperStep(data["mapping"].(map[string]string)),
	}
}
