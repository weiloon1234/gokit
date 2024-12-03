package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Confirmed validates that the value matches the field with the same name plus "_confirmation"
// Params: []
func Confirmed(value interface{}, params []string, c *gin.Context) bool {
	if len(params) < 1 {
		return false // Parameter specifying the field to confirm is required
	}

	// Get the field name from the params
	fieldName := params[0]
	confirmationFieldName := fmt.Sprintf("%s_confirmation", fieldName)

	// Retrieve the original and confirmation values from the request context
	originalValue := c.PostForm(fieldName)
	confirmationValue := c.PostForm(confirmationFieldName)

	return originalValue == confirmationValue
}
