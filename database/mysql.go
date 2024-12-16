package database

import (
	"context"
	"database/sql"
	"fmt"

	"time"

	"entgo.io/ent/dialect"

	entSQL "entgo.io/ent/dialect/sql"
	"github.com/hashicorp/go-multierror"
	"github.com/weiloon1234/gokit/config"
	"github.com/weiloon1234/gokit/ent"
	"github.com/weiloon1234/gokit/ent/hook"
	"github.com/weiloon1234/gokit/ent/migrate"
)

var (
	dbClient       *ent.Client
	sqlDB          *sql.DB // Save the raw database connection for later management
	GlobalDBConfig *config.DBConfig
)

// Init initializes the Ent client and connects to the database.
func Init(config *config.DBConfig) error {
	SetGlobalDBConfig(config)
	// Open database connection
	var err error
	sqlDB, err = sql.Open("mysql", config.GetDSN())
	fmt.Printf(config.GetDSN())
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

	// Initialize the Ent client
	dbClient = ent.NewClient(ent.Driver(entDriver))

	// Run the schema migration with context timeout
	migrationCtx, cancelMigration := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelMigration()
	if err := dbClient.Schema.Create(
		migrationCtx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		return fmt.Errorf("failed to create schema resources: %w", err)
	}

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

// GetDBClient retrieves the Ent client.
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
