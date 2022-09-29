package transformer

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer/steps"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type TransformerMessage struct {
	queue.MessageData
	data any
}

func (m TransformerMessage) GetData() (any, error) {
	return json.Marshal(m.data)
}

// Transformer is the management point of the transformer.
type Transformer struct {
	name  string
	steps []steps.Step
}

func (t Transformer) Name() string {
	return t.name
}

func (t Transformer) Steps() []steps.Step {
	return t.steps
}

func NewTransformer(name string, steps []steps.Step) Transformer {
	return Transformer{
		name:  name,
		steps: steps,
	}
}


func (transformer Transformer) runSteps(input queue.Message) (any, error) {
	output, err := input.GetData()
	if err != nil {
		return nil, err
	}
	for _, step := range transformer.Steps() {
		output, err = step.Process(output)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func (transformer Transformer) sendMessage(output any, outQueue queue.Queue) error {
	message := TransformerMessage{
		queue.NewMessageData(),
		output,
	}
	return outQueue.Add(message)
}

func (transformer Transformer) transform(inMessage queue.Message, outQueue queue.Queue) error {
	output, err := transformer.runSteps(inMessage)
	if err != nil {
		return err
	}

	return transformer.sendMessage(output, outQueue)
}

// StartTransformer is the controller for the transformer.
// It will start the transformer and will listen for messages, as well as waiting for cxt to be done.
func (transformer Transformer) StartTransformer(cxt context.Context, inQueue queue.Queue, outQueue queue.Queue) {
	for {
		select {
		case <-cxt.Done():
			return
		case inMessage := <-inQueue.Chan():
			err := transformer.transform(inMessage, outQueue)
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			time.Sleep(10 * time.Second)
		}
	}
}
