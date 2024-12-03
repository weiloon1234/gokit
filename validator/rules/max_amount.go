package rules

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MaxAmount checks if the value is less than or equal to the specified maximum value
// Params: [maxValue]
func MaxAmount(value interface{}, params []string, c *gin.Context) bool {
	if len(params) == 0 {
		return false // Maximum value is required as a parameter
	}

	// Convert the parameter to a float
	maxValue, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return false
	}

	// Convert the value to a float
	strValue := fmt.Sprintf("%v", value)
	floatValue, err := strconv.ParseFloat(strValue, 64)
	if err != nil || floatValue < 0 {
		return false // Value must be a valid number and non-negative
	}

	return floatValue <= maxValue
}
