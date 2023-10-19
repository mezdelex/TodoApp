package main

import (
	"encoding/json"
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"todoapp.com/application/services"
	"todoapp.com/configuration"
	"todoapp.com/domain/models"
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

	var config models.Config
	error = configuration.LoadCfg(&config)
	if error != nil {
		log.Fatal("Error decoding loaded configuration file.")
	}

	encodedKey, error := os.ReadFile(config.PrivateKeyPath)
	if error != nil {
		log.Fatal("Error accessing OPENSSH public key.")
	}

	privateKey, error := jwt.ParseRSAPrivateKeyFromPEM(encodedKey)
	if error != nil {
		log.Fatal("Error parsing OPENSSH public key.")
	}

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	db := connectors.Postgre{}.Connect()

	// repositories
	todosRepository := repositories.NewTodosRepository(db)
	usersRepository := repositories.NewUsersRepository(db)

	// services
	todosService := services.NewTodosService(todosRepository)
	usersService := services.NewUsersService(usersRepository)
	loginService := services.NewLoginService(usersRepository, &config)

	// controllers
	todosController := controllers.NewTodosController(todosService)
	usersController := controllers.NewUsersController(usersService)
	loginController := controllers.NewLoginController(loginService)

	// unsecured routes
	loginController.Route(app.Group("/api", logger.New()))

	// secured routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    privateKey.Public(),
		},
	}))
	api := app.Group("/api", logger.New())
	todosController.Route(api)
	usersController.Route(api)

	app.Listen(":3000")
}
