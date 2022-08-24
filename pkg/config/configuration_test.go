package config

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/consumer/twitch"
	"config_con/pkg/utils/environment"
	"reflect"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	environment.Settings.ConfigFilePath = "test_data/test_config.yaml"
	tests := []struct {
		name    string
		want    YamlConfiguration
		wantErr bool
	}{
		{
			name: "ReadConfiguration",
			want: YamlConfiguration{
				Consumers: []consumer.ConsumerConfig{
					{
						TwitchEventConfigs: []twitch.TwitchEventConsumer{
							{
								Name:        "test_consumer",
								EventSecret: "test_consumer_secret",
								Url:         "test_consumer_url",
							},
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
