package transformer

import "config_con/pkg/pipe/transformer/steps"

type Step interface {
	Process(any) (any, error)
}

// StepConfig is the holder for the specific step configuration.
type StepConfig struct {
	HashMapperSteps []steps.MapperStep `yaml:"hashMapperSteps"`
}

func (stepConfig StepConfig) GetStepMap() map[string]Step {
	stepMap := make(map[string]Step)
	for _, hashMapperStep := range stepConfig.HashMapperSteps {
		stepMap[hashMapperStep.Name] = hashMapperStep
	}
	return stepMap
}

// TransformerStepConfig holds what steps belong to the transformer.
type TransformerStepConfig struct {
	Name  string   `yaml:"name"`
	Steps []string `yaml:"steps"`
}

// TransformerConfig holds the configuration for the transformers.
type TransformerConfig struct {
	Transformers []TransformerStepConfig `yaml:"transformers"`
	Steps        StepConfig            `yaml:"steps"`
}

func (config TransformerConfig) GetTransformerMap() map[string]Transformer {
	transformerMap := make(map[string]Transformer)
	steps := config.Steps.GetStepMap()
	for _, transformer := range config.Transformers {
		transformerSteps := []Step{}

		for _, stepName := range transformer.Steps {
			transformerSteps = append(transformerSteps, steps[stepName])
		}
		transformerMap[transformer.Name] = Transformer{
			Name:  transformer.Name,
			Steps: transformerSteps,
		}
	}
	return transformerMap
}

type Transformer struct {
	Name  string
	Steps []Step
}
