package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/globalcontext"
)

func GlobalContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a unique request ID
		requestID := globalcontext.GenerateRequestID()
		c.Set("requestID", requestID)

		// Create a base context with admin_path information
		ctx := context.WithValue(context.Background(), "gin_context", c)

		// Store the context in the global context manager
		globalcontext.SetContext(requestID, ctx)

		// Proceed with the request
		c.Next()

		// Clear the context after the request ends
		globalcontext.ClearContext(requestID)
	}
}
