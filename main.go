package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"todoapp.com/infrastructure/connectors/postgre"
	"todoapp.com/presentation/controllers/todos"
)

func main() {
	postgre.Connect()

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	api := app.Group("/api", logger.New())

	todos.SetupTodosController(api)

	app.Listen(":3000")
}
