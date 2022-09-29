package transformer

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer/steps"
	"config_con/pkg/utils/test"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeStep struct {
	name     string
	throwErr bool
}

func (s fakeStep) Process(data any) (any, error) {
	if s.throwErr {
		return nil, fmt.Errorf("error")
	}
	return data.(string) + " test", nil
}

func TestTransformer_RunSteps(t *testing.T) {
	type fields struct {
		Name  string
		Steps []steps.Step
	}
	type args struct {
		input queue.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "TestTransformer_RunSteps",
			fields: fields{
				Name: "testTransformer",
				Steps: []steps.Step{
					fakeStep{
						name: "fakeStep1",
					},
					fakeStep{
						name: "fakeStep2",
					},
				},
			},
			args: args{
				input: test.NewFakeMessage("test"),
			},
			want:    "test test test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := NewTransformer(tt.fields.Name, tt.fields.Steps)
			got, err := test.runSteps(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transformer.RunSteps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transformer.RunSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformer_transform(t *testing.T) {
	fakeQueue := queue.NewLocalQueue(1)
	defer fakeQueue.Close()
	step := fakeStep{
		name: "fakeStep",
	}

	transformer := Transformer{
		name:  "testTransformer",
		steps: []steps.Step{step},
	}

	err := transformer.transform(test.NewFakeMessage("test"), fakeQueue)
	assert.NoError(t, err)

	data := <-fakeQueue.Chan()

	out, err := data.GetData()
	assert.NoError(t, err)
	assert.Equal(t, "\"test test\"", string(out.([]byte)))

	transformer = Transformer{
		name: "testTransformer",
		steps: []steps.Step{
			fakeStep{
				name:     "fakeStep1",
				throwErr: true,
			},
		},
	}
	err = transformer.transform(test.NewFakeMessage("test"), fakeQueue)
	assert.Error(t, err)
}
