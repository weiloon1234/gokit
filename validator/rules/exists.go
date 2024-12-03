package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/database"
)

// Exists validates that the value exists in the specified database table and column
func Exists(value interface{}, params []string, c *gin.Context) bool {
	if len(params) < 2 {
		return false // Invalid usage
	}
	tableName := params[0]
	columnName := params[1]

	var count int64
	database.GetDB().Table(tableName).Where(fmt.Sprintf("%s = ?", columnName), value).Count(&count)

	return count > 0
}
