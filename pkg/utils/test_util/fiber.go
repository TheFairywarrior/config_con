package testutil

import "encoding/json"


type FakeFiberContext struct {
	Body []byte
	Headers map[string]string
}

func (f *FakeFiberContext) BodyParser(v interface{}) error {
	return json.Unmarshal(f.Body, v)
}
