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
					"single":      "single",
				},
			},
			want: map[string]any{
				"general": map[string]any{
					"kenobi": map[string]any(nil),
				},
				"single": map[string]any(nil),
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

func TestMapperStep_AddData(t *testing.T) {
	type fields struct {
		Name      string
		MapConfig map[string]string
	}
	type args struct {
		data    map[string]any
		newData map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]any
	}{
		{
			name: "TestMapperStep_AddData",
			fields: fields{
				Name: "TestMapperStep_AddData",
				MapConfig: map[string]string{
					"hello.there": "general.kenobi",
					"single":      "single",
				},
			},
			args: args{
				data: map[string]any{
					"hello": map[string]any{
						"there": "hello there value",
					},
					"single": "single key value",
				},
				newData: map[string]any{
					"general": map[string]any{
						"kenobi": map[string]any(nil),
					},
					"single": map[string]any(nil),
				},
			},
			want: map[string]any{
				"general": map[string]any{
					"kenobi": "hello there value",
				},
				"single": "single key value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := MapperStep{
				Name:      tt.fields.Name,
				MapConfig: tt.fields.MapConfig,
			}
			if got := mapper.AddData(tt.args.data, tt.args.newData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapperStep.AddData() = %v, want %v", got, tt.want)
			}
		})
	}
}
