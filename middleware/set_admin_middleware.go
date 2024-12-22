package middleware

import (
	"github.com/gin-gonic/gin"
)

// AdminMiddleware sets admin context for audit logs.
func SetAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: add the real logic to get the admin object(struct) from the request
		admin := map[string]string{"id": "1", "name": "admin_name"}
		c.Set("admin", admin)
		c.Set("is_admin", true)

		c.Next()
	}
}
