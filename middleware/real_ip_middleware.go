package middleware

import (
	"github.com/weiloon1234/gokit/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func RealIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the "X-Forwarded-For" header
		xForwardedFor := c.GetHeader("X-Forwarded-For")

		// Split the header if it contains multiple IPs
		if xForwardedFor != "" {
			ips := strings.Split(xForwardedFor, ",")
			for _, ip := range ips {
				ip = strings.TrimSpace(ip)
				if utils.IsValidIP(ip) {
					c.Set("RealIP", ip)
					c.Request.RemoteAddr = ip // Optionally overwrite RemoteAddr
					return
				}
			}
		}

		// Fall back to "X-Real-IP"
		xRealIP := c.GetHeader("X-Real-IP")
		if xRealIP != "" && utils.IsValidIP(xRealIP) {
			c.Set("RealIP", xRealIP)
			c.Request.RemoteAddr = xRealIP // Optionally overwrite RemoteAddr
			return
		}

		// Fall back to ClientIP
		clientIP := c.ClientIP()
		if utils.IsValidIP(clientIP) {
			c.Set("RealIP", clientIP)
			c.Request.RemoteAddr = clientIP // Optionally overwrite RemoteAddr
			return
		}

		// If no valid IP found, leave the IP unset or handle it as needed
		c.Set("RealIP", "")
	}
}
