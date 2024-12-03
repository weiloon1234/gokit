package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

// IsUsername checks if the value is an alphanumeric string or underscore with 6 to 32 characters
func IsUsername(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)
	return regexp.MustCompile(`^[A-Za-z0-9_]{6,32}$`).MatchString(strValue)
}
