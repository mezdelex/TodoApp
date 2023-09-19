package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	todosService "todoapp.com/application/services/todos"
	"todoapp.com/infrastructure/connectors/postgre"
	"todoapp.com/infrastructure/environments"
	todosRepository "todoapp.com/infrastructure/repositories/todos"
	todosController "todoapp.com/presentation/controllers/todos"
)

func main() {
	error := environments.LoadEnv()
	if error != nil {
		log.Fatal("Error loading .env file.")
	}

	db := postgre.Connect()

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	api := app.Group("/api", logger.New())

	// repositories
	todosRepository := todosRepository.NewTodosRepository(db)

	// services
	todosService := todosService.NewTodosService(todosRepository)

	// controllers
	todosController := todosController.NewTodosController(todosService)

	// routes
	todosController.Route(api)

	app.Listen(":3000")
}
