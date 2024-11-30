package rules

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

func Alphanumeric(value interface{}, param string, c *gin.Context) bool {
	alphaNumRegex := `^[a-zA-Z0-9]+$`
	re := regexp.MustCompile(alphaNumRegex)
	return re.MatchString(fmt.Sprintf("%v", value))
}
