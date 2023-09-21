package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"todoapp.com/application/services"
	"todoapp.com/infrastructure/connectors"
	"todoapp.com/infrastructure/environments"
	"todoapp.com/infrastructure/repositories"
	"todoapp.com/presentation/controllers"
)

func main() {
	error := environments.LoadEnv()
	if error != nil {
		log.Fatal("Error loading .env file.")
	}

	db := connectors.Postgre{}.Connect()

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	api := app.Group("/api", logger.New())

	// repositories
	todosRepository := repositories.NewTodosRepository(db)
	_ = repositories.NewUsersRepository(db)

	// services
	todosService := services.NewTodosService(todosRepository)

	// controllers
	todosController := controllers.NewTodosController(todosService)

	// routes
	todosController.Route(api)

	app.Listen(":3000")
}
