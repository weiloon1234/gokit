package rules

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Min validates that the value's length is at least the given parameter
func Min(value interface{}, param string, c *gin.Context) bool {
	valLen := len(fmt.Sprintf("%v", value))
	minVal, err := strconv.Atoi(param)
	if err != nil {
		return false
	}
	return valLen >= minVal
}
