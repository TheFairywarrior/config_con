package transformer

import (
	"fmt"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer/steps"
)

// TransformerConfig holds the configuration for the transformers.
type TransformerConfig struct {
	Transformers []TransformerStepConfig `yaml:"transformers"`
	Steps        StepConfig              `yaml:"steps"`
}

func (config TransformerConfig) GetTransformerMap() (map[string]transformer.Transformer, error) {
	transformerMap := make(map[string]transformer.Transformer)
	stepsMaps := config.Steps.GetStepMap()
	for _, transformerConfig := range config.Transformers {
		transformerSteps := []steps.Step{}

		for _, stepName := range transformerConfig.Steps {
			step, ok := stepsMaps[stepName]
			if !ok {
				return nil, fmt.Errorf("transformer '%s' had an error: step \"%s\" not found", transformerConfig.Name, stepName)
			}

			transformerSteps = append(transformerSteps, step)
		}
		transformerMap[transformerConfig.Name] = transformer.NewTransformer(transformerConfig.Name, transformerSteps)
	}
	return transformerMap, nil
}
