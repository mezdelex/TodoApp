package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"todoapp.com/application/services"
	"todoapp.com/infrastructure/connectors/postgre"
	"todoapp.com/infrastructure/repositories"
	"todoapp.com/presentation/controllers"
)

func main() {
	db := postgre.Connect()

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	api := app.Group("/api", logger.New())

	// repositories
	todosRepository := repositories.NewTodosRepository(db)

	// services
	todosService := services.NewTodosService(todosRepository)

	// controllers
	todosController := controllers.NewTodosController(todosService)

	// routes
	todosController.Route(api)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("DB_PORT")))
}
