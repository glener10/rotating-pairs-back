package main

import (
	"fmt"
	"os"

	"github.com/glener10/rotating-pairs-back/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := LoadEnvironmentVariables(); err != nil {
		fmt.Println("Error to load environment variables: ", err)
		return
	}

	routes.HandlerRoutes()
}

func LoadEnvironmentVariables() error {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return nil
	}
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
