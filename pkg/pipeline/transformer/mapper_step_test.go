package transformer

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
				mapConfig: tt.fields.MapConfig,
			}
			if got := mapper.build(); !reflect.DeepEqual(got, tt.want) {
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
		name    string
		fields  fields
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "TestMapperStep_AddData success",
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
			wantErr: false,
		},
		{
			name: "TestMapperStep_AddData success",
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
						"broken": "hello there value",
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
			want:    map[string]any(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := MapperStep{
				mapConfig: tt.fields.MapConfig,
			}
			got, err := mapper.AddData(tt.args.data, tt.args.newData)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapperStep.AddData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapperStep.AddData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapperStep_Process(t *testing.T) {
	type fields struct {
		Name      string
		MapConfig map[string]string
	}
	type args struct {
		data any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "TestMapperStep_Process success",
			fields: fields{
				Name: "TestMapperStep_Process",
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
			},
			want: map[string]any{
				"general": map[string]any{
					"kenobi": "hello there value",
				},
				"single": "single key value",
			},
			wantErr: false,
		},
		{
			name: "TestMapperStep_Process failure",
			fields: fields{
				Name: "TestMapperStep_Process",
				MapConfig: map[string]string{
					"hello.there": "general.kenobi",
					"single":      "single",
				},
			},
			args: args{
				data: map[string]any{
					"hello": map[string]any{
						"broken": "hello there value",
					},
					"single": "single key value",
				},
			},
			want:    map[string]any(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := MapperStep{
				mapConfig: tt.fields.MapConfig,
			}
			got, err := mapper.Process(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapperStep.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapperStep.Process() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMapperStep(t *testing.T) {
	type args struct {
		name      string
		mapConfig map[string]string
	}
	tests := []struct {
		name string
		args args
		want MapperStep
	}{
		{
			name: "TestNewMapperStep success",
			args: args{
				name: "TestNewMapperStep",
				mapConfig: map[string]string{
					"hello.there": "general.kenobi",
					"single":      "single",
				},
			},
			want: MapperStep{
				mapConfig: map[string]string{
					"hello.there": "general.kenobi",
					"single":      "single",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMapperStep(tt.args.mapConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapperStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
