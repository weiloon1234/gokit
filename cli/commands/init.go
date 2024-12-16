package commands

import (
	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit/seeds"
)

// InitCommands adds all core commands to the root command
func InitCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(MigrateCmd)
	rootCmd.AddCommand(SeederCmd)

	RegisterSeeder("country_seeder", seeds.CountrySeeder)
}
