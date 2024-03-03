package Middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MethodsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		if method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		allowedMethods := []string{"GET", "POST"}
		found := false
		for _, m := range allowedMethods {
			if m == method {
				found = true
				break
			}
		}
		if !found {
			statusCode := http.StatusForbidden
			c.AbortWithStatusJSON(statusCode, gin.H{"error": "Method not allowed", "statusCode": statusCode})
		} else {
			c.Next()
		}
	}
}
