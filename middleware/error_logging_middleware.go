package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ErrorLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request

		// Check for errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				log.Printf("Critical error: %v", e.Err)
			}
		}
	}
}
