package Middlewares

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func OriginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Host
		origin2 := c.Request.Header.Get("Origin")

		originAllowed := false
		allowedUrls := getAllowedURLs()
		for _, allowedOrigin := range allowedUrls {
			if origin == allowedOrigin || allowedOrigin == "*" {
				originAllowed = true
				break
			}
		}

		if !originAllowed {
			messageOriginNotAllowed := "Origin '" + origin + "' not allowed, second origin: " + origin2
			c.AbortWithStatusJSON(403, gin.H{"error": messageOriginNotAllowed})
		} else {
			c.Next()
		}
	}
}

func getAllowedURLs() []string {
	allowedURLsString := os.Getenv("ALLOW_URLS")
	if allowedURLsString == "" {
		return nil
	}
	allowedUrlsString := strings.Split(allowedURLsString, "|")
	return allowedUrlsString
}
