package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

// IsNumber checks if the value consists only of digits
func IsNumber(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)
	return regexp.MustCompile(`^\d+$`).MatchString(strValue)
}
