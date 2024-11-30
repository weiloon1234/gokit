package commands

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var seederRegistry = map[string]func(){}

// RegisterSeeder allows external projects to add seeders to the registry
func RegisterSeeder(name string, seederFunc func()) {
	seederRegistry[name] = seederFunc
}

func availableSeeders() []string {
	keys := make([]string, 0, len(seederRegistry))
	for k := range seederRegistry {
		keys = append(keys, k)
	}
	return keys
}

var SeederCmd = &cobra.Command{
	Use:   "seed [seeder_name]",
	Short: "Run a specific database seeder",
	Long:  "Run a specific database seeder by name",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		seederName := args[0]

		// Find the seeder function in the registry
		seederFunc, found := seederRegistry[seederName]
		if !found {
			log.Errorf("Seeder '%s' not found. Available seeders: %v", seederName, availableSeeders())
			return errors.New(fmt.Sprintf("Seeder '%s' not found. Available seeders: %v", seederName, availableSeeders()))
		}

		// Run the seeder function
		log.Infof("Running seeder: %s", seederName)
		seederFunc()
		return nil
	},
}
