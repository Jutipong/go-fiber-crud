package routes

import (
	"fiber-crud/app/address/controller"
	"fiber-crud/app/address/repository"
	"fiber-crud/app/address/service"
	"fiber-crud/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	repository := repository.NewRepository(config.Db())
	service := service.NewService(repository)
	controller := controller.NewController(service)
	auth := app.Group("/auth")
	auth.Get("/login", controller.Login)
	auth.Post("/create", controller.Create)
}
