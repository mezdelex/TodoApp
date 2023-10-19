package configuration

import (
	"encoding/json"
	"log"
	"os"

	"todoapp.com/domain/models"
)

func LoadCfg(config *models.Config) error {
	configFile, error := os.ReadFile("config.json")
	if error != nil {
		log.Fatal("Configuration could not be loaded.")
	}

	return json.Unmarshal(configFile, config)
}
