package rules

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Max validates that the value's length or numeric value does not exceed the given parameter
func Max(value interface{}, params []string, c *gin.Context) bool {
	if len(params) == 0 {
		return false // Parameter required for Max rule
	}
	maxVal, err := strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	switch v := value.(type) {
	case int, int64, float32, float64:
		valFloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
		return valFloat <= float64(maxVal)
	case string:
		return len(v) <= maxVal
	default:
		valLen := len(fmt.Sprintf("%v", value))
		return valLen <= maxVal
	}
}
