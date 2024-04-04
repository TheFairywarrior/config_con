package config

import (
	"reflect"
	"testing"
)

type testConfig struct {
	name      string
	something string
}

func (c *testConfig) Load() (any, error) {
	return nil, nil
}

func (c *testConfig) Validate() error {
	return nil
}

func (c *testConfig) Name() string {
	return c.name
}

func (c *testConfig) Type() string {
	return "consumer"
}

func testConfigConstructor(config map[string]any) Configuration {
	return &testConfig{
		name:      config["name"].(string),
		something: config["something"].(string),
	}
}

func TestLoad(t *testing.T) {
	Register("test", testConfigConstructor)
	type args struct {
		configs map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "Test Load",
			args: args{
				configs: map[string]any{
					"consumer": map[string]any{
						"test": []map[string]any{
							{
								"name":      "test",
								"something": "something",
							},
						},
					},
					"transformer": map[string]any{
						"test": []map[string]any{
							{
								"name":      "test",
								"something": "something",
							},
						},
					},
					"publisher": map[string]any{
						"test": []map[string]any{
							{
								"name":      "test",
								"something": "something",
							},
						},
					},
				},
			},
			want: Config{
				consumerConfigs: map[string]Configuration{
					"test": &testConfig{
						name:      "test",
						something: "something",
					},
				},
				transformerConfigs: map[string]Configuration{
					"test": &testConfig{
						name:      "test",
						something: "something",
					},
				},
				publisherConfigs: map[string]Configuration{
					"test": &testConfig{
						name:      "test",
						something: "something",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.configs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
