package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Load(app *fiber.App) {
	api := app.Group("/x-mens")
	mutant(api)
	stats(api)
}

func mutant(api fiber.Router) {
	api.Post("/mutant", Health)
}

func stats(api fiber.Router) {
	api.Get("/stats", Health)
}