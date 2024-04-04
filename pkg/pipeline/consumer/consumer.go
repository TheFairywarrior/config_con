package consumer

import "context"


type Consumer interface {
	Consume(ctx context.Context) error
}


var consumerInstances = map[string]Consumer{}
var consumers = map[string]func(map[string]any) Consumer{}


func Register(name string, constructor func(map[string]any) Consumer) {
	consumers[name] = constructor
}


