package consumer

import "context"


type ComsumedOut interface {
	Add(context.Context, map[string]string) error
	Chan() chan map[string]string
}

type Consumer interface {
	Consume(context.Context, chan ComsumedOut)
}
