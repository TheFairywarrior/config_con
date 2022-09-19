package test

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type FakeFiberContext struct {
	OutBody       any
	CurrentStatus int
	RequestBody   []byte
	Headers       map[string]string
}

func (f *FakeFiberContext) BodyParser(v interface{}) error {
	return json.Unmarshal(f.RequestBody, v)
}

func (f *FakeFiberContext) GetReqHeaders() map[string]string {
	return f.Headers
}

func (f *FakeFiberContext) Status(status int) *fiber.Ctx {
	f.CurrentStatus = status
	return nil
}

func (f *FakeFiberContext) JSON(v interface{}) error {
	f.OutBody = v
	return nil
}

func (f *FakeFiberContext) Send(body []byte) error {
	f.OutBody = body
	return nil
}

func (f *FakeFiberContext) Body() []byte {
	return f.RequestBody
}
