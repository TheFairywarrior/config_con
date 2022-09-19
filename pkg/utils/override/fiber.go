package override

import "github.com/gofiber/fiber/v2"

type FiberContext interface {
	BodyParser(interface{}) error
	GetReqHeaders() map[string]string
	Status(status int) *fiber.Ctx
	JSON(interface{}) error
	Send([]byte) error
	Body() []byte
}
