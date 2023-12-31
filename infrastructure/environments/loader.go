package environments

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

const DefaultEnvPath = "/.env"

func LoadEnv() error {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error generating .env dir.")
	}
	dir := filepath.Join(filepath.Dir(f), "../..", DefaultEnvPath)

	err := godotenv.Load(dir)

	return err
}
