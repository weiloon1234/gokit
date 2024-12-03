package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// DatetimeISO8601 checks if the value is a valid ISO8601 datetime string with a timezone
// Params: [] (no additional parameters required)
func DatetimeISO8601(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)

	// Attempt to parse the value as an ISO8601 date-time string with timezone
	_, err := time.Parse(time.RFC3339, strValue)

	// Return true if parsing was successful, otherwise false
	return err == nil
}
