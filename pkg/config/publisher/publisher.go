package publisher

import (
	fileConfig "config_con/pkg/config/publisher/file"

	"config_con/pkg/pipe/publisher"
	"config_con/pkg/pipe/publisher/file"
)

// PublisherConfig is the configuration for the publisher.
type PublisherConfig struct {
	FilePublisherConfig []fileConfig.FilePublisherConfig `yaml:"filePublishers"`
}

func (publisherConfig PublisherConfig) GetPublisherMap() map[string]publisher.Publisher {
	publisherMap := make(map[string]publisher.Publisher)
	for _, filePublisherConfig := range publisherConfig.FilePublisherConfig {
		filePublisher := file.NewFilePublisher(filePublisherConfig.Name, filePublisherConfig.FilePath, filePublisherConfig.FileMode)
		publisherMap[filePublisherConfig.Name] = filePublisher
	}
	return publisherMap
}
