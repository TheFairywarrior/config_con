package transformer

import (
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer/steps"
	"reflect"
	"testing"
)

func TestTransformerConfig_BuildTransformerMap(t *testing.T) {
	type fields struct {
		Transformers []TransformerStepConfig
		Steps        StepConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]transformer.Transformer
		wantErr bool
	}{
		{
			name: "Test successful config",
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
					HashMapperSteps: []MapperStepConfig{
						{
							Name: "from1toone",
							MapConfig: map[string]string{
								"thing.one": "value",
							},
						},
					},
				},
			},
			want: map[string]transformer.Transformer{
				"transformer1": transformer.NewTransformer("transformer1", []steps.Step{
					steps.NewMapperStep("from1toone", map[string]string{
						"thing.one": "value",
					}),
				}),
			},
			wantErr: false,
		},
		{
			name: "Test non existent step",
			fields: fields{
				Transformers: []TransformerStepConfig{
					{
						Name: "transformer1",
						Steps: []string{
							"nonexistent",
						},
					},
				},
				Steps: StepConfig{
					HashMapperSteps: []MapperStepConfig{
						{
							Name: "from1toone",
							MapConfig: map[string]string{
								"thing.one": "value",
							},
						},
					},
				},
			},
			want:    map[string]transformer.Transformer(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := TransformerConfig{
				Transformers: tt.fields.Transformers,
				Steps:        tt.fields.Steps,
			}
			got, err := config.GetTransformerMap()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformerConfig.BuildTransformerMap() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformerConfig.BuildTransformerMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
