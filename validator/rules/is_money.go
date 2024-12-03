package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// IsMoney checks if the value is a float >= 0
func IsMoney(value interface{}, params []string, c *gin.Context) bool {
	strValue := fmt.Sprintf("%v", value)
	floatValue, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return false
	}
	return floatValue >= 0
}
