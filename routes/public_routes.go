package routes

import (
	"fiber-crud/features/address/controller"
	"fiber-crud/features/address/repository"
	"fiber-crud/features/address/service"
	"fiber-crud/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	repository := repository.NewRepository(config.Db())
	service := service.NewService(repository)
	controller := controller.NewController(service)
	auth := app.Group("/address")
	auth.Get("/inquiry/:id", controller.Inquiry)
	// auth.Post("/create", controller.Create)
	// auth.Post("/testdecimal", controller.TestDecimal)
	// auth.Put("/update", controller.Update)
	// auth.Delete("/delete/:id", controller.Delete)
}
