package publisher

import "config_con/pkg/pipe/publisher/file"

// Publisher interface used as a base for all publishers.
type Publisher interface {
	// Publish is used to publish the data to the publisher.
	Publish(data []byte) error
}

// PublisherConfig is the configuration for the publisher.
type PublisherConfig struct {
	FilePublisher []file.FilePublisher `yaml:"filePublishers"`
}

func (publisher PublisherConfig) GetPublisherMap() map[string]Publisher {
	publisherMap := make(map[string]Publisher)
	for _, filePublisher := range publisher.FilePublisher {
		publisherMap[filePublisher.Name] = filePublisher
	}
	return publisherMap
}
