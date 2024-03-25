package publisher

import (
	fileConfig "github.com/thefairywarrior/config_con/pkg/config/publisher/file"
	"github.com/thefairywarrior/config_con/pkg/pipe/publisher"
	"github.com/thefairywarrior/config_con/pkg/pipe/publisher/file"
	"reflect"
	"testing"
)

func TestPublisherConfig_GetPublisherMap(t *testing.T) {
	type fields struct {
		FilePublisherConfig []fileConfig.FilePublisherConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]publisher.Publisher
	}{
		{
			name: "test",
			fields: fields{
				FilePublisherConfig: []fileConfig.FilePublisherConfig{
					{
						Name:     "test",
						FilePath: "test",
						FileMode: 0644,
					},
				},
			},
			want: map[string]publisher.Publisher{
				"test": file.NewFilePublisher("test", "test", 0644),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publisher := PublisherConfig{
				FilePublisherConfig: tt.fields.FilePublisherConfig,
			}
			if got := publisher.GetPublisherMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublisherConfig.GetPublisherMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
