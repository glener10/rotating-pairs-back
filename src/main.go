package main

import (
	"fmt"

	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/glener10/rotating-pairs-back/src/routes"
)

func main() {
	if err := Utils.LoadEnvironmentVariables(); err != nil {
		fmt.Println("Error to load environment variables: ", err)
		return
	}

	routes.HandlerRoutes()
}
