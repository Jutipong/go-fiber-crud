package main

import (
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/enum"
	"fiber-crud/pkg/utils"
	"fiber-crud/routes"

	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/shopspring/decimal"
)

func init() {
	config.InitialConfig()
	config.InitialDB()
}

type Result struct {
	Name    string
	Address string
}

func main() {
	// decimal.MarshalJSONWithoutQuotes = true
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*"}))

	// Middleware
	app.Use(middleware.Logger)
	// app.Use(middleware.Authorization())
	// genSQL()

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
