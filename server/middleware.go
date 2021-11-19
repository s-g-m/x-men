package server

import (
	"net/http"
	"x-mers/config"

	"github.com/gofiber/fiber/v2"
)

func validateAccess(c *fiber.Ctx) error {
	if config.ApiKey() == c.Get("x-api-key") {
		return c.Next()
	}
	return c.Status(http.StatusUnauthorized).Send([]byte("invalid api key"))
}
