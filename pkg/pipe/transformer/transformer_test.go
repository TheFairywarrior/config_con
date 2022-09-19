package transformer

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer/steps"
	"config_con/pkg/utils/test"
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
		name    string
		fields  fields
		want    map[string]Transformer
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
			want: map[string]Transformer(nil),
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

type fakeStep struct {
	name string
}

func (s fakeStep) Process(data any) (any, error) {
	return data.(string) + " test", nil
}

func TestTransformer_RunSteps(t *testing.T) {
	type fields struct {
		Name  string
		Steps []Step
	}
	type args struct {
		input queue.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "TestTransformer_RunSteps",
			fields: fields{
				Name: "testTransformer",
				Steps: []Step{
					fakeStep{
						name: "fakeStep1",
					},
					fakeStep{
						name: "fakeStep2",
					},
				},
			},
			args: args{
				input: test.NewFakeMessage("test"),
			},
			want:    "test test test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformer := Transformer{
				Name:  tt.fields.Name,
				Steps: tt.fields.Steps,
			}
			got, err := transformer.runSteps(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transformer.RunSteps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transformer.RunSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
