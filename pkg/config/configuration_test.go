package config

import (
	"config_con/pkg/pipe"
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/consumer/twitch"
	"config_con/pkg/pipe/publisher"
	"config_con/pkg/pipe/publisher/file"
	"config_con/pkg/pipe/transformer"
	"config_con/pkg/pipe/transformer/steps"
	"config_con/pkg/utils/environment"
	"context"
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
				Consumers: consumer.ConsumerConfig{
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

func TestYamlConfiguration_CreatePipelines(t *testing.T) {
	type fields struct {
		Consumers    consumer.ConsumerConfig
		Transformers transformer.TransformerConfig
		Publishers   publisher.PublisherConfig
		Pipelines    []pipe.PipeConfig
	}
	type args struct {
		cxt context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]pipe.Pipe
		wantErr bool
	}{
		{
			name: "CreatePipelines_Test",
			fields: fields{
				Consumers: consumer.ConsumerConfig{
					TwitchEventConfigs: []twitch.TwitchEventConsumer{
						{
							Name:        "test_consumer",
							EventSecret: "test_consumer_secret",
							Url:         "test_consumer_url",
						},
					},
				},
				Transformers: transformer.TransformerConfig{
					Transformers: []transformer.TransformerStepConfig{
						{
							Name: "test_transformer",
							Steps: []string{
								"test_step",
							},
						},
					},
					Steps: transformer.StepConfig{
						HashMapperSteps: []steps.MapperStep{
							{
								Name: "test_step",
								MapConfig: map[string]string{
									"test_key": "test_value",
								},
							},
						},
					},
				},
				Publishers: publisher.PublisherConfig{
					FilePublisher: []file.FilePublisher{
						{
							Name:     "test_publisher",
							FilePath: "test_path",
							FileMode: 0644,
						},
					},
				},
				Pipelines: []pipe.PipeConfig{
					{
						Name:        "test_pipeline",
						Consumer:    "test_consumer",
						Transformer: "test_transformer",
						Publisher:   "test_publisher",
					},
				},
			},
			args: args{
				cxt: context.Background(),
			},
			want: map[string]pipe.Pipe{
				"test_pipeline": pipe.NewPipe(
					context.Background(),
					twitch.TwitchEventConsumer{
						Name:        "test_consumer",
						EventSecret: "test_consumer_secret",
						Url:         "test_consumer_url",
					},
					transformer.Transformer{
						Name: "test_transformer",
						Steps: []transformer.Step{
							steps.MapperStep{
								Name: "test_step",
								MapConfig: map[string]string{
									"test_key": "test_value",
								},
							},
						},
					},
					file.FilePublisher{
						Name:     "test_publisher",
						FilePath: "test_path",
						FileMode: 0644,
					},
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := YamlConfiguration{
				Consumers:    tt.fields.Consumers,
				Transformers: tt.fields.Transformers,
				Publishers:   tt.fields.Publishers,
				Pipelines:    tt.fields.Pipelines,
			}
			got, err := config.CreatePipelines(tt.args.cxt)
			if (err != nil) != tt.wantErr {
				t.Errorf("YamlConfiguration.CreatePipelines() error = %v, wantErr %v", err, tt.wantErr)
			}
			gotMap := got["test_pipeline"]
			wantMap := tt.want["test_pipeline"]

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("YamlConfiguration.CreatePipelines() = %v, want %v", got, tt.want)
			}
		})
	}
}
