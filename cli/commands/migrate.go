package commands

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" // Replace with your DB driver
	_ "github.com/golang-migrate/migrate/v4/source/file"    // For file source migrations
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	core "github.com/weiloon1234/gokit"
	"os/exec"
	"strconv"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate [action] [optional arguments]",
	Short: "Run database migrations (e.g., create, up, down, force)",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		action := args[0]
		config := core.GetSharedConfig()

		dsn := config.DBConfig.DSN
		if dsn == "" {
			return errors.New("DB_DSN not set in environment variables or .env file")
		}

		switch action {
		case "create":
			if len(args) < 2 {
				return errors.New("please provide a name for the migration")
			}
			migrationName := args[1]
			return createMigration(migrationName)

		case "up":
			return runMigration(dsn, "up")

		case "down":
			return runMigration(dsn, "down")

		case "force":
			if len(args) < 2 {
				return errors.New("please provide a version number to force")
			}
			version, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid version number: %v", err)
			}
			return forceMigration(dsn, version)

		default:
			return fmt.Errorf("unknown action: %s. Available actions: create, up, down, force", action)
		}
	},
}

func createMigration(name string) error {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "database/migrations", "-seq", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Failed to create migration: %v", err)
		fmt.Println(string(output))
		return err
	}
	fmt.Printf("Migration '%s' created successfully\n", name)
	return nil
}

func runMigration(dsn string, direction string) error {
	m, err := migrate.New("file://database/migrations", dsn)
	if err != nil {
		return fmt.Errorf("migration setup failed: %v", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migration up failed: %v", err)
		}
		fmt.Println("Migration up completed successfully")

	case "down":
		if err := m.Down(); err != nil {
			return fmt.Errorf("migration down failed: %v", err)
		}
		fmt.Println("Migration down completed successfully")
	}

	return nil
}

func forceMigration(dsn string, version int) error {
	m, err := migrate.New("file://database/migrations", dsn)
	if err != nil {
		return fmt.Errorf("migration setup failed: %v", err)
	}

	if err := m.Force(version); err != nil {
		return fmt.Errorf("migration force failed: %v", err)
	}
	fmt.Printf("Migration forced to version %d successfully\n", version)
	return nil
}
