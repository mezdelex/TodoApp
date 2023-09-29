package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"todoapp.com/domain/models"
)

func LoadCfg(config *models.Config) error {
	configFile, error := os.ReadFile("config.json")
	if error != nil {
		log.Fatal("Configuration could not be loaded.")
	}

	error = json.Unmarshal(configFile, config)
	fmt.Println("Estoy en la config")
	fmt.Println(config)

	return error
}
