package Utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() error {
	pathToEnv := "../../.env"
	if _, err := os.Stat(pathToEnv); os.IsNotExist(err) {
		return nil
	}
	if err := godotenv.Load(pathToEnv); err != nil {
		return err
	}
	return nil
}
