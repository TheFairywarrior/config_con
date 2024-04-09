package runner

import (
	"github.com/thefairywarrior/config_con/pkg/pipeline/consumer"
	"github.com/thefairywarrior/config_con/pkg/pipeline/publisher"
	"github.com/thefairywarrior/config_con/pkg/pipeline/transformer"
)

type Pipeline struct {
	consumer       consumer.Consumer
	transformSteps []transformer.TransformStep
	publishers     map[string]publisher.Publisher
}
