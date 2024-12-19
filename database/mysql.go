package database

import (
	"context"
	"database/sql"
	"fmt"

	"reflect"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-multierror"
	"github.com/weiloon1234/gokit-base-entity/ent/hook"
	"github.com/weiloon1234/gokit/config"
	"github.com/weiloon1234/gokit/logger"
)

var (
	sqlDB          *sql.DB      // Save the raw database connection for later management
	entClient      *interface{} // Use an interface to avoid direct dependency on the ent package
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

func SetEntClient(client interface{}) error {
	entClient = &client

	// Apply soft-delete hook
	if err := hook.AddSoftDeleteFilter(client); err != nil {
		return fmt.Errorf("Failed to add soft-delete filter: %v", err)
	}

	return nil
}

func GetEntClient() interface{} {
	if entClient == nil {
		return nil // Return nil if the client hasn't been set
	}
	return *entClient
}

func WithTransaction(ctx *gin.Context, fn func(tx interface{}) error) error {
	clientVal := reflect.ValueOf(GetEntClient())
	txMethod := clientVal.MethodByName("Tx")
	if !txMethod.IsValid() {
		return fmt.Errorf("dbClient does not have a 'Tx' method")
	}

	// Start the transaction
	txResults := txMethod.Call([]reflect.Value{reflect.ValueOf(ctx)})
	if len(txResults) == 0 || txResults[0].IsNil() {
		return fmt.Errorf("failed to start transaction")
	}
	tx := txResults[0]

	// Defer rollback in case of panic or early return
	defer func() {
		if r := recover(); r != nil {
			tx.MethodByName("Rollback").Call(nil)
			logger.GetLogger().Printf("Transaction rolled back due to panic: %v", r)
		}
	}()

	// Run the transactional function
	err := fn(tx.Interface())
	if err != nil {
		tx.MethodByName("Rollback").Call(nil)
		return err
	}

	// Commit the transaction
	commitResults := tx.MethodByName("Commit").Call(nil)
	if len(commitResults) > 0 && !commitResults[0].IsNil() {
		return commitResults[0].Interface().(error)
	}

	return nil
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
