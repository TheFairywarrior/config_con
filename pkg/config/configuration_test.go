package config

import (
	"context"
	"reflect"
	"testing"

	"github.com/thefairywarrior/config_con/pkg/pipe/consumer"
	consumerConfig "github.com/thefairywarrior/config_con/pkg/config/consumer"
	"github.com/thefairywarrior/config_con/pkg/config/publisher"
	fileConfig "github.com/thefairywarrior/config_con/pkg/config/publisher/file"
	transformerConfig "github.com/thefairywarrior/config_con/pkg/config/transformer"
	"github.com/thefairywarrior/config_con/pkg/pipe"
	"github.com/thefairywarrior/config_con/pkg/pipe/publisher/file"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer"
	"github.com/thefairywarrior/config_con/pkg/pipe/transformer/steps"

	"github.com/thefairywarrior/config_con/pkg/utils/environment"
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
				Consumers: consumerConfig.ConsumerConfig{
					TwitchEventConfigs: []consumerConfig.TwitchEventConfig{
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
		Consumers    consumerConfig.ConsumerConfig
		Transformers transformerConfig.TransformerConfig
		Publishers   publisher.PublisherConfig
		Pipelines    []pipe.PipeConfig
	}
	type args struct {
		ctx context.Context
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
				Consumers: consumerConfig.ConsumerConfig{
					TwitchEventConfigs: []consumerConfig.TwitchEventConfig{
						{
							Name:        "test_consumer",
							EventSecret: "test_consumer_secret",
							Url:         "test_consumer_url",
						},
					},
				},
				Transformers: transformerConfig.TransformerConfig{
					Transformers: []transformerConfig.TransformerStepConfig{
						{
							Name: "test_transformer",
							Steps: []string{
								"test_step",
							},
						},
					},
					Steps: transformerConfig.StepConfig{
						HashMapperSteps: []transformerConfig.MapperStepConfig{
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
					FilePublisherConfig: []fileConfig.FilePublisherConfig{
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
				ctx: context.Background(),
			},
			want: map[string]pipe.Pipe{
				"test_pipeline": pipe.NewPipe(
					context.Background(),
					consumer.NewTwitchEventConsumer("test_consumer", "test_consumer_secret", "test_consumer_url"),
					transformer.NewTransformer("test_transformer", []steps.Step{
						steps.NewMapperStep("test_step", map[string]string{
							"test_key": "test_value",
						}),
					}),
					file.NewFilePublisher("test_publisher", "test_path", 0644),
				),
			},
			wantErr: false,
		},
		{
			name: "CreatePipelines_Test_Invalid_Pipeline",
			fields: fields{
				Consumers: consumerConfig.ConsumerConfig{
					TwitchEventConfigs: []consumerConfig.TwitchEventConfig{
						{
							Name:        "doesnt_exist",
							EventSecret: "test_consumer_secret",
							Url:         "test_consumer_url",
						},
					},
				},
				Transformers: transformerConfig.TransformerConfig{
					Transformers: []transformerConfig.TransformerStepConfig{
						{
							Name: "test_transformer",
							Steps: []string{
								"test_step",
							},
						},
					},
					Steps: transformerConfig.StepConfig{
						HashMapperSteps: []transformerConfig.MapperStepConfig{
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
					FilePublisherConfig: []fileConfig.FilePublisherConfig{
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
				ctx: context.Background(),
			},
			want:    map[string]pipe.Pipe{},
			wantErr: true,
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
			got, err := config.CreatePipelines(tt.args.ctx)
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
