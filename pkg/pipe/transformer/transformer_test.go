package transformer

import (
	"config_con/pkg/pipe/queue"
	"config_con/pkg/pipe/transformer/steps"
	"config_con/pkg/utils/test"
	"reflect"
	"testing"
)

type fakeStep struct {
	name string
}

func (s fakeStep) Process(data any) (any, error) {
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
