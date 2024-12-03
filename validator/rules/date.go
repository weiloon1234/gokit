package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Date checks if the value is a valid date in various common formats
func Date(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)

	// Define common date formats
	formats := []string{
		"2006-01-02",          // yyyy-mm-dd
		"2006-01-02 15:04",    // yyyy-mm-dd hh:mm
		"2006-01-02 15:04:05", // yyyy-mm-dd hh:mm:ss
		"02-Jan-2006",         // dd-Mmm-yyyy
		"02/01/2006",          // dd/mm/yyyy
	}

	// Try parsing the value with each of the common formats
	for _, format := range formats {
		if _, err := time.Parse(format, strValue); err == nil {
			return true // Successfully parsed, valid date
		}
	}

	return false // Could not parse with any formats
}
