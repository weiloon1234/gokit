package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

// IsPassword2 checks if the value is exactly 6 digits
func IsPassword2(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)
	return regexp.MustCompile(`^[0-9]{6}$`).MatchString(strValue)
}
