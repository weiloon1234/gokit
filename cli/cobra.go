package cli

import (
	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit/cli/commands"
	"log"
)

// rootCmd is the root of the CLI
var rootCmd = &cobra.Command{
	Use:   "gokit",
	Short: "CLI for gokit tasks",
	Long:  "CLI for running various gokit tasks such as migrations and seeding",
}

// RootCmd exposes the root command for extensions
func RootCmd() *cobra.Command {
	return rootCmd
}

// Init initializes core commands and adds additional commands
func Init(additionalCmds []*cobra.Command) {
	commands.InitCommands(rootCmd)
	rootCmd.AddCommand(additionalCmds...)
}

// ExecuteCLI runs the CLI with the configured commands
func ExecuteCLI() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing CLI: %v", err)
	}
}
