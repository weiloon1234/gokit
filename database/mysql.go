package database

import (
	"context"
	"database/sql"
	"fmt"

	"reflect"
	"time"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"

	entSQL "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-multierror"
	"github.com/weiloon1234/gokit-base-entity/ent/hook"
	"github.com/weiloon1234/gokit/config"
	"github.com/weiloon1234/gokit/logger"
)

var (
	dbClient       interface{}
	sqlDB          *sql.DB // Save the raw database connection for later management
	GlobalDBConfig *config.DBConfig
)

// Init initializes the Ent client and connects to the database.
func Init(config *config.DBConfig, entPackage interface{}) error {
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

	// Validate that the entPackage contains the expected components
	entVal := reflect.ValueOf(entPackage)
	clientField := entVal.FieldByName("NewClient")
	if !clientField.IsValid() {
		return fmt.Errorf("entPackage does not have a 'NewClient' function")
	}

	// Call the NewClient function to create the ent.Client
	newClientFunc := clientField.Interface().(func() interface{})
	dbClient = newClientFunc()

	// Wrap the database connection with Ent's SQL driver
	entDriver := entSQL.OpenDB(dialect.MySQL, sqlDB)

	// Use reflection to set the driver on the ent.Client
	clientVal := reflect.ValueOf(dbClient)
	driverMethod := clientVal.MethodByName("Driver")
	if !driverMethod.IsValid() {
		return fmt.Errorf("entClient does not have a 'Driver' method")
	}

	driverMethod.Call([]reflect.Value{reflect.ValueOf(entDriver)})

	// Run schema migrations (if Schema exists)
	schemaField := entVal.FieldByName("Schema")
	if !schemaField.IsValid() {
		return fmt.Errorf("entPackage does not have a 'Schema' function")
	}

	schemaVal := schemaField.Interface()
	schema := reflect.ValueOf(schemaVal)

	createMethod := schema.MethodByName("Create")
	if !createMethod.IsValid() {
		return fmt.Errorf("entClient does not have a 'Create' method")
	}

	migrationCtx, cancelMigration := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelMigration()
	createResults := createMethod.Call([]reflect.Value{
		reflect.ValueOf(migrationCtx),
		reflect.ValueOf(true),
		reflect.ValueOf(true),
	})

	if len(createResults) > 0 && !createResults[0].IsNil() {
		return fmt.Errorf("failed to create schema resources: %w", createResults[0].Interface().(error))
	}

	// Add the soft-delete filter
	if err := hook.AddSoftDeleteFilter(dbClient); err != nil {
		return fmt.Errorf("failed to add soft-delete filter: %w", err)
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

// GetDBClient retrieves the Ent client as an interface{}.
func GetDBClient() interface{} {
	return dbClient
}

func WithTransaction(ctx *gin.Context, fn func(tx interface{}) error) error {
	clientVal := reflect.ValueOf(dbClient)
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
