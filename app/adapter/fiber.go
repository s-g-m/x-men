package adapter

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type fiberAdapter struct {
	fiberCtx *fiber.Ctx
}

func NewFiberAdapter(c *fiber.Ctx) HttpAdapter {
	return &fiberAdapter{fiberCtx: c}
}

func (f fiberAdapter) SendResponse(status int, response interface{}) {
	if reflect.TypeOf(response).Kind() == reflect.String {
		response = map[string]interface{}{
			"message": response,
		}
	}
	f.fiberCtx.Status(status).JSON(response)
}

func (f fiberAdapter) GetBody() (body []byte) {
	body = f.fiberCtx.Body()
	return
}
