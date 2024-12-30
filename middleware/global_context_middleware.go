package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/globalcontext"
	"github.com/weiloon1234/gokit/logger"
)

func GlobalContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a unique request ID
		requestID := globalcontext.GenerateRequestID()
		c.Set("requestID", requestID)

		// Create a base context with admin_path information
		ctx := context.WithValue(context.Background(), globalcontext.GinCtxKey, c)

		// Store the context in the global context manager
		globalcontext.SetContext(requestID, ctx)

		logger.GetLogger().Info("RequestID: " + requestID)

		// Proceed with the request
		c.Next()

		// Clear the context after the request ends
		globalcontext.ClearContext(requestID)
	}
}
