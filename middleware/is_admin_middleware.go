package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

type contextKey string

const IsAdminKey contextKey = "is_admin"

func IsAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/admin") {
			ctx := context.WithValue(c.Request.Context(), IsAdminKey, true)

			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}
	}
}

func IsInAdminRoute(ctx context.Context) bool {
	isAdmin, ok := ctx.Value(IsAdminKey).(bool)
	return ok && isAdmin
}
