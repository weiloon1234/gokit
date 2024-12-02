package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/logger"
)

func ErrorLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request

		// Check for errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logger.GetLogger().Printf("Critical error: %v", e.Err)
			}
		}
	}
}
