package transformer

import (
	"config_con/pkg/pipe/transformer/steps"
	"reflect"
	"testing"
)

func TestStepConfig_GetStepMap(t *testing.T) {
	type fields struct {
		HashMapperSteps []steps.MapperStep
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]Step
	}{
		{
			name: "TestStepConfig_GetStepMap",
			fields: fields{
				HashMapperSteps: []steps.MapperStep{
					{
						Name: "hashMapperStep",
						MapConfig: map[string]string{
							"key": "value",
						},
					},
					{
						Name: "hashMapperSte2",
						MapConfig: map[string]string{
							"key2": "value2",
						},
					},
				},
			},
			want: map[string]Step{
				"hashMapperStep": steps.MapperStep{
					Name: "hashMapperStep",
					MapConfig: map[string]string{
						"key": "value",
					},
				},
				"hashMapperSte2": steps.MapperStep{
					Name: "hashMapperSte2",
					MapConfig: map[string]string{
						"key2": "value2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stepConfig := StepConfig{
				HashMapperSteps: tt.fields.HashMapperSteps,
			}
			if got := stepConfig.GetStepMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StepConfig.GetStepMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformerConfig_BuildTransformerMap(t *testing.T) {
	type fields struct {
		Transformers []TransformerStepConfig
		Steps        StepConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]Transformer
	}{
		{
			name: "TestTransformerConfig_BuildTransformerMap",
			fields: fields{
				Transformers: []TransformerStepConfig{
					{
						Name: "transformer1",
						Steps: []string{
							"from1toone",
						},
					},
				},
				Steps: StepConfig{
					HashMapperSteps: []steps.MapperStep{
						{
							Name: "from1toone",
							MapConfig: map[string]string{
								"thing.one": "value",
							},
						},
					},
				},
			},
			want: map[string]Transformer{
				"transformer1": {
					Name: "transformer1",
					Steps: []Step{
						steps.MapperStep{
							Name: "from1toone",
							MapConfig: map[string]string{
								"thing.one": "value",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := TransformerConfig{
				Transformers: tt.fields.Transformers,
				Steps:        tt.fields.Steps,
			}
			if got := config.GetTransformerMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformerConfig.BuildTransformerMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
