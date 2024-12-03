package rules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/database"
)

// Unique validates that the value is unique in the specified database table and column
// Params: [tableName, columnName, excludeColumnName (optional), excludeValue (optional)]
func Unique(value interface{}, params []string, c *gin.Context) bool {
	if len(params) < 2 {
		return false // At least table name and column name are required
	}

	tableName := params[0]
	columnName := params[1]

	var count int64
	query := database.GetDB().Table(tableName).Where(fmt.Sprintf("%s = ?", columnName), value)

	// Iterate through the remaining parameters in pairs to apply exclusion conditions
	for i := 2; i+1 < len(params); i += 2 {
		excludeColumn := params[i]
		excludeValue := params[i+1]
		query = query.Not(fmt.Sprintf("%s = ?", excludeColumn), excludeValue)
	}

	query.Count(&count)

	// If count is greater than zero, value is not unique
	return count == 0
}
