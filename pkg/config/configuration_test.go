package config

import (
	"config_con/pkg/utils/shortcuts"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	os.Setenv("CONFIG_FILE_PATH", "test_data/test_config.yaml")
	asdf, _ := os.Getwd()
	fmt.Println(asdf)
	tests := []struct {
		name    string
		want    shortcuts.Map
		wantErr bool
	}{
		{
			name: "TestReadConfiguration",
			want: shortcuts.Map{
				"consumers": []shortcuts.Map{
					{
						"name": "consumer1",
						"type": "api",
						"configuration": shortcuts.Map{
							"eventSecret":        "secret",
							"eventWebHookRoute": "webhook",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
