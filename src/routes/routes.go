package routes

import (
	"github.com/gin-gonic/gin"
	CombinationGenerationCounterRoute "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/routes"
)

func HandlerRoutes() {
	r := gin.Default()
	CombinationGenerationCounterRoute.HandlerRoute(r)
	r.Run()
}
