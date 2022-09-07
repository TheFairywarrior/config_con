package steps

import (
	"reflect"
	"testing"
)

func TestMapperStep_Build(t *testing.T) {
	type fields struct {
		Name      string
		MapConfig map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]any
	}{
		{
			name: "TestMapperStep_Build",
			fields: fields{
				Name: "TestMapperStep_Build",
				MapConfig: map[string]string{
					"hello.there": "general.kenobi",
					"single": "single",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := MapperStep{
				Name:      tt.fields.Name,
				MapConfig: tt.fields.MapConfig,
			}
			if got := mapper.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapperStep.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
