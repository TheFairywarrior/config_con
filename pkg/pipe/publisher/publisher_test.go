package publisher

import (
	"config_con/pkg/pipe/publisher/file"
	"reflect"
	"testing"
)

func TestPublisherConfig_GetPublisherMap(t *testing.T) {
	type fields struct {
		FilePublisher []file.FilePublisher
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]Publisher
	}{
		{
			name: "test",
			fields: fields{
				FilePublisher: []file.FilePublisher{
					{
						Name:     "test",
						FilePath: "test",
						FileMode: 0644,
					},
				},
			},
			want: map[string]Publisher{
				"test": file.FilePublisher{
					Name:     "test",
					FilePath: "test",
					FileMode: 0644,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publisher := PublisherConfig{
				FilePublisher: tt.fields.FilePublisher,
			}
			if got := publisher.GetPublisherMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublisherConfig.GetPublisherMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
