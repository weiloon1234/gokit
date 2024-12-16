package rules

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/database"
)

// Exists validates that the value exists in the specified database table and column
func Exists(value interface{}, params []string, c *gin.Context) bool {
	if len(params) < 2 {
		return false // Insufficient parameters: tableName and columnName are required
	}

	// Extract table and column name from params
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

	// Execute the query
	var count int
	err := sqlDB.QueryRowContext(c.Request.Context(), query, value).Scan(&count)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return false
	}

	// Return true if at least one record exists
	return count > 0
}
