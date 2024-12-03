package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

// IsContactNumber checks if the value is a number with 8 to 12 digits
func IsContactNumber(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)
	return regexp.MustCompile(`^[0-9]{8,12}$`).MatchString(strValue)
}
