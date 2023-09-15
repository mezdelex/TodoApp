package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"todoapp.com/application/services"
	"todoapp.com/infrastructure/connectors/postgre"
	"todoapp.com/infrastructure/repositories"
	"todoapp.com/presentation/controllers"
)

func main() {
	postgre.Connect()

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	api := app.Group("/api", logger.New())

	// repositories
	todosRepository := repositories.NewTodosRepository(postgre.DB)

	// services
	todosService := services.NewTodosService(todosRepository)

	// controllers
	todosController := controllers.NewTodosController(todosService)

	// routes
	todosController.Route(api)

	app.Listen(":3000")
}
