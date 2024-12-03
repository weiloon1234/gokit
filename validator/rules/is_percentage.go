package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

// IsPercentage checks if the value is a valid percentage (integer or float)
func IsPercentage(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)
	return regexp.MustCompile(`^\d+(?:\.\d+)?$`).MatchString(strValue)
}
