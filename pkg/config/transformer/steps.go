package transformer

import (
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer/steps"
)

// StepConfig is the holder for the specific step configuration.
type StepConfig struct {
	HashMapperSteps []MapperStepConfig `yaml:"hashMapperSteps"`
}

func (stepConfig StepConfig) GetStepMap() map[string]steps.Step {
	stepMap := make(map[string]steps.Step)
	for _, hashMapperStep := range stepConfig.HashMapperSteps {
		hashMapper := steps.NewMapperStep(hashMapperStep.Name, hashMapperStep.MapConfig)
		stepMap[hashMapperStep.Name] = hashMapper
	}
	return stepMap
}

// TransformerStepConfig holds what steps belong to the transformer.
type TransformerStepConfig struct {
	Name  string   `yaml:"name"`
	Steps []string `yaml:"steps"`
}

type MapperStepConfig struct {
	Name      string            `yaml:"name"`
	MapConfig map[string]string `yaml:"mapConfig"`
}
