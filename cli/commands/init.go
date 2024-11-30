package commands

import "github.com/spf13/cobra"

// InitCommands adds all core commands to the root command
func InitCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(MigrateCmd)
	rootCmd.AddCommand(SeederCmd)
}
