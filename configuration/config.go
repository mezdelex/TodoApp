package configuration

import (
	"encoding/json"
	"log"
	"os"

	"todoapp.com/domain/models"
)

func LoadCfg(config *models.Config) error {
	configFile, error := os.Open("config.json")
	if error != nil {
		log.Fatal("Configuration could not be loaded.")
	}
	defer configFile.Close()

	error = json.NewDecoder(configFile).Decode(config)

	return error
}
