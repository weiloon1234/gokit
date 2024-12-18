package database

import (
	"context"
	"database/sql"
	"fmt"

	"time"

	"entgo.io/ent/dialect"

	entSQL "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-multierror"
	"github.com/weiloon1234/gokit/config"
	"github.com/weiloon1234/gokit/ent"
	"github.com/weiloon1234/gokit/ent/hook"
	"github.com/weiloon1234/gokit/logger"
)

var (
	dbClient       *ent.Client
	sqlDB          *sql.DB // Save the raw database connection for later management
	GlobalDBConfig *config.DBConfig
)

// Init initializes the Ent client and connects to the database.
func Init(config *config.DBConfig, clientFactory func(driver ent.Driver) *ent.Client) error {
	var err error

	SetGlobalDBConfig(config)
	// Open database connection
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

	// Wrap the database connection with Ent's SQL driver
	entDriver := entSQL.OpenDB(dialect.MySQL, sqlDB)

	// Initialize the ent client using the clientFactory
	dbClient = clientFactory(entDriver) // Initialize the ent client using the clientFactory

	// Add the soft-delete filter
	hook.AddSoftDeleteFilter(dbClient)

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

// GetDBClient returns the Ent client, ensuring the project schema is included.
func GetDBClient() *ent.Client {
	return dbClient
}

func GetSQLDB() *sql.DB {
	return sqlDB
}

// CloseDB safely closes the database connection and the Ent client.
func CloseDB() error {
	var errs error

	// Close Ent client
	if dbClient != nil {
		if err := dbClient.Close(); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("failed to close Ent client: %w", err))
		}
	}

	// Close raw SQL connection
	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("failed to close SQL connection: %w", err))
		}
	}

	// Return combined errors or nil if no errors occurred
	return errs
}

func WithTransaction(ctx *gin.Context, fn func(tx *ent.Tx) error) error {
	// Start the transaction
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return err // Return error if the transaction can't be started
	}

	// Defer rollback in case of panic or early return
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback transaction on panic
			logger.GetLogger().Printf("Transaction rolled back due to panic: %v", r)
		}
	}()

	// Run the transactional function
	err = fn(tx)
	if err != nil {
		tx.Rollback() // Rollback transaction on error
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err // Return error if commit fails
	}

	return nil
}
