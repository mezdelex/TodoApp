package connectors

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todoapp.com/domain/models"
)

type Postgre struct{}

func (_ Postgre) Connect() *gorm.DB {
	port, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Couldn't connect to the database.")
	}
	db.AutoMigrate(&models.Todo{}, &models.User{})

	fmt.Println("Connected to the DB.")
	return db
}
