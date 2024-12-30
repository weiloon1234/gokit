package globalcontext

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ctxStore sync.Map

// GenerateRequestID generates a unique request ID.
func GenerateRequestID() string {
	return uuid.New().String()
}

// SetContext stores the context for the given request ID.
func SetContext(requestID string, ctx context.Context) {
	ctxStore.Store(requestID, ctx)
}

// GetContext retrieves the context for the given request ID.
func GetContext(requestID string) context.Context {
	if ctx, ok := ctxStore.Load(requestID); ok {
		return ctx.(context.Context)
	}
	return context.Background()
}

// ClearContext removes the context for the given request ID.
func ClearContext(requestID string) {
	ctxStore.Delete(requestID)
}

func GetGinContext(ctx context.Context) *gin.Context {
	if ginCtx, ok := ctx.Value("gin_context").(*gin.Context); ok {
		return ginCtx
	}
	return nil
}
