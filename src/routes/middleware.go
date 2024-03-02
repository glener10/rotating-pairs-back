package routes

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Host

		originAllowed := false
		allowedUrls := GetAllowedURLs()
		for _, allowedOrigin := range allowedUrls {
			if origin == allowedOrigin || allowedOrigin == "*" {
				originAllowed = true
				break
			}
		}

		if !originAllowed {
			c.AbortWithStatusJSON(403, gin.H{"error": "Origin not allowed"})
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", strings.Join(allowedUrls, ","))
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func GetAllowedURLs() []string {
	allowedURLsString := os.Getenv("ALLOW_URLS")
	if allowedURLsString == "" {
		return nil
	}
	allowedUrlsString := strings.Split(allowedURLsString, "|")
	return allowedUrlsString
}
