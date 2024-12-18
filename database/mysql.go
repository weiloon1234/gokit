package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"entgo.io/ent"
	entSQL "entgo.io/ent/dialect/sql"
	"github.com/weiloon1234/gokit/config"
)

var (
	dbClient *ent.Client
	sqlDB    *sql.DB
)

// Init initializes the Ent client and connects to the database.
func Init(config *config.DBConfig, clientFactory func(driver ent.Driver) *ent.Client) error {
	var err error

	// Open the database connection
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

	// Wrap the database connection with Ent's driver
	entDriver := entSQL.OpenDB(ent.DriverNameMySQL, sqlDB)

	// Initialize the Ent client using the provided factory function
	dbClient = clientFactory(entDriver)
	if dbClient == nil {
		return fmt.Errorf("failed to initialize ent client")
	}

	return nil
}

// GetDBClient retrieves the Ent client.
func GetDBClient() *ent.Client {
	return dbClient
}
