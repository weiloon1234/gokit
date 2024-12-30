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
		var isAdmin bool
		if strings.HasPrefix(c.Request.URL.Path, "/api/admin") {
			isAdmin = true
		} else {
			isAdmin = false
		}

		ctx := context.WithValue(c.Request.Context(), IsAdminKey, isAdmin)
		c.Set(string(IsAdminKey), isAdmin)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func IsInAdminRoute(ctx context.Context) bool {
	isAdmin, ok := ctx.Value(IsAdminKey).(bool)
	return ok && isAdmin
}
