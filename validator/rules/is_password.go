package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

// IsPassword checks if the value has between 6 and 24 characters
func IsPassword(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)
	return regexp.MustCompile(`^.{6,24}$`).MatchString(strValue)
}
