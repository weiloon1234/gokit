package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entSQL "entgo.io/ent/dialect/sql"
	"github.com/weiloon1234/gokit/config"
	"github.com/weiloon1234/gokit/ent"
	"github.com/weiloon1234/gokit/ent/hook"
)

var (
	dbClient       *ent.Client
	sqlDB          *sql.DB // Save the raw database connection for later management
	GlobalDBConfig *config.DBConfig
)

// Init initializes the Ent client and connects to the database.
func Init(config *config.DBConfig) error {
	SetGlobalDBConfig(config)

	// Open the database connection using the standard library
	var err error
	sqlDB, err = sql.Open("mysql", config.GetDSN())
	if err != nil {
		return fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	// Test the database connection
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	// Wrap the database connection with Ent's SQL driver
	entDriver := entSQL.OpenDB(dialect.MySQL, sqlDB)

	// Initialize the Ent client
	dbClient = ent.NewClient(ent.Driver(entDriver))

	// Run the schema migration
	if err := dbClient.Schema.Create(context.Background()); err != nil {
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
func CloseDB() {
	// Close the Ent client first
	if dbClient != nil {
		if err := dbClient.Close(); err != nil {
			log.Printf("Error while closing Ent client: %v", err)
		} else {
			log.Println("Ent client closed.")
		}
	}

	// Close the raw database connection
	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error while closing database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}
