package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// DateFormat checks if the value matches the specified date format
// Params: [format]
func DateFormat(value interface{}, params []string, c *gin.Context) bool {
	if len(params) == 0 {
		return false // The format is required as a parameter
	}

	strValue := fmt.Sprintf("%v", value)
	dateFormat := params[0]

	// Try to parse the value with the provided format
	_, err := time.Parse(dateFormat, strValue)
	return err == nil // Return true if parsing was successful
}
