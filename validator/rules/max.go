package rules

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Max validates that the value's length does not exceed the given parameter
func Max(value interface{}, param string, c *gin.Context) bool {
	valLen := len(fmt.Sprintf("%v", value))
	maxVal, err := strconv.Atoi(param)
	if err != nil {
		return false
	}
	return valLen <= maxVal
}
