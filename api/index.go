package handler

import (
	"fmt"

	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/glener10/rotating-pairs-back/src/routes"
)

func Handler() {
	if err := Utils.LoadEnvironmentVariables("../.env"); err != nil {
		fmt.Println("Error to load environment variables: ", err)
		return
	}

	r := routes.HandlerRoutes()
	routes.Listening(r)
}
