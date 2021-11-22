package routes

import (
	"x-men/app/modules/dna"
	"x-men/app/modules/statistic"
	"x-men/app/repository"
	"x-men/server/handlers"

	"github.com/gofiber/fiber/v2"
)

func Load(app *fiber.App) {
	repository := repository.NewRepository()

	statisticService := statistic.NewService(repository)
	statisticController := statistic.NewController(statisticService)

	dnaService := dna.NewService(repository)
	dnaController := dna.NewController(dnaService)

	api := app.Group("/x-men")

	stats(statisticController, api)
	mutant(dnaController, api)
}

func mutant(controller dna.Controller, api fiber.Router) {
	handler := handlers.NewDnaHandlers(controller)
	api.Post("/mutant", handler.Mutant)
}

func stats(controller statistic.Controller, api fiber.Router) {
	handler := handlers.NewStatisticHandlers(controller)
	api.Get("/stats", handler.Stats)
}
