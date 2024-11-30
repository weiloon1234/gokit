package rules

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Email validates that the value is a valid email address
func Email(value interface{}, param string, c *gin.Context) bool {
	emailRegex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(fmt.Sprintf("%v", value))
}
