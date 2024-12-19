package database

import (
	"context"
	"database/sql"
	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-multierror"
	"github.com/weiloon1234/gokit/config"
)

var (
	sqlDB          *sql.DB // Save the raw database connection for later management
	GlobalDBConfig *config.DBConfig
)

// Init initializes the Ent client and connects to the database.
func Init(config *config.DBConfig) error {
	SetGlobalDBConfig(config)
	// Open database connection
	var err error
	sqlDB, err = sql.Open("mysql", config.GetDSN())
	if err != nil {
		return fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	// Test the database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	return nil
}

// SetGlobalDBConfig sets the global database configuration.
func SetGlobalDBConfig(config *config.DBConfig) {
	GlobalDBConfig = config
}

// GetGlobalDBConfig retrieves the global database configuration.
func GetGlobalDBConfig() *config.DBConfig {
	return GlobalDBConfig
}

func GetSQLDB() *sql.DB {
	return sqlDB
}

// CloseDB safely closes the database connection and the Ent client.
func CloseDB() error {
	var errs error

	// Close raw SQL connection
	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("failed to close SQL connection: %w", err))
		}
	}

	// Return combined errors or nil if no errors occurred
	return errs
}
