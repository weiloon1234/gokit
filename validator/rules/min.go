package rules

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Min validates that the value's length or numeric value is not less than the given parameter
func Min(value interface{}, params []string, c *gin.Context) bool {
	if len(params) == 0 {
		return false // Parameter required for Min rule
	}
	minVal, err := strconv.Atoi(params[0])
	if err != nil {
		return false
	}

	switch v := value.(type) {
	case int, int64, float32, float64:
		valFloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
		return valFloat >= float64(minVal)
	case string:
		return len(v) >= minVal
	default:
		valLen := len(fmt.Sprintf("%v", value))
		return valLen >= minVal
	}
}
