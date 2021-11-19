package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) (err error) {
	message := []byte(`{"message":"The service x-mens is active"}`)
	c.Status(http.StatusOK).Send(message)
	return
}
