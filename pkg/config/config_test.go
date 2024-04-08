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
	Register("testCon", testConfigConstructor)
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
						"testCon": []map[string]any{
							{
								"name":      "test1",
								"something": "something",
							},
						},
					},
					"transformer": map[string]any{
						"testCon": []map[string]any{
							{
								"name":      "test2",
								"something": "something",
							},
						},
					},
					"publisher": map[string]any{
						"testCon": []map[string]any{
							{
								"name":      "test3",
								"something": "something",
							},
						},
					},
				},
			},
			want: Config{
				consumerConfigs: map[string]Configuration{
					"test1": &testConfig{
						name:      "test1",
						something: "something",
					},
				},
				transformerConfigs: map[string]Configuration{
					"test2": &testConfig{
						name:      "test2",
						something: "something",
					},
				},
				publisherConfigs: map[string]Configuration{
					"test3": &testConfig{
						name:      "test3",
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
