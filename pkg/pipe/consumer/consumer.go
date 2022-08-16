package consumer

import "context"


type ComsumedOut interface {

}

type Consumer interface {
	Consume(context.Context, chan ComsumedOut)
}