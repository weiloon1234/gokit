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

	// Get the raw SQL database connection from the global sqlDB variable
	sqlDB := database.GetSQLDB()
	if sqlDB == nil {
		fmt.Printf("Database connection is not initialized\n")
		return false
	}

	// Construct the raw SQL query
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", tableName, columnName)
	args := []interface{}{value}

	// Iterate through the remaining parameters in pairs to apply exclusion conditions
	for i := 2; i+1 < len(params); i += 2 {
		excludeColumn := params[i]
		excludeValue := params[i+1]

		// Append exclusion condition
		query += fmt.Sprintf(" AND %s != ?", excludeColumn)
		args = append(args, excludeValue)
	}

	var count int64
	err := sqlDB.QueryRowContext(c.Request.Context(), query, args...).Scan(&count)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return false
	}

	// If count is greater than zero, value is not unique
	return count == 0
}
