package main

import (
	"fiber-crud/middleware"
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/enum"
	"fiber-crud/pkg/utils"
	"fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.InitialConfig()
	config.InitialDB()
}

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*"}))

	// Middleware
	app.Use(middleware.Logger)
	app.Use(middleware.Authorization())

	// Routes
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	_configEnv := config.Server()
	if _configEnv.Env_Mode == enum.ModeDebug {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
