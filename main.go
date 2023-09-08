package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	postgre_connector "todoapp.com/infrastructure"
	todoscontroller "todoapp.com/presentation/controllers"
)

func main() {
	postgre_connector.Connect()

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	api := app.Group("/api", logger.New())

	todoscontroller.SetupTodosController(api)

	app.Listen(":3000")
}
