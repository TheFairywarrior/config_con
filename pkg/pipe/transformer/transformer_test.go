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
						MapConfig: map[string]string {
							"key": "value",
						},
					},
					{
						Name: "hashMapperSte2",
						MapConfig: map[string]string {
							"key2": "value2",
						},
					},
				},
			},
			want: map[string]Step{
				"hashMapperStep": steps.MapperStep{
					Name: "hashMapperStep",
					MapConfig: map[string]string {
						"key": "value",
					},
				},
				"hashMapperSte2": steps.MapperStep{
					Name: "hashMapperSte2",
					MapConfig: map[string]string {
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
