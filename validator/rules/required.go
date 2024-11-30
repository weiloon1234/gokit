package rules

import (
	"github.com/gin-gonic/gin"
)

// Required validates that the value is not empty or nil
func Required(value interface{}, param string, c *gin.Context) bool {
	return value != nil && value != ""
}
