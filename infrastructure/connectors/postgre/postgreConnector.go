package postgre

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todoapp.com/domain/models/todo"
)

var DB *gorm.DB

func Connect() {
	var error error

	error = godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file.")
	}

	port, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Panic("Couldn't connect to the database.")
	}
	DB.AutoMigrate(&todo.Todo{})

	fmt.Println("Connected to the DB.")
}
