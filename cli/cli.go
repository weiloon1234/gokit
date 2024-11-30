package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute initializes and starts the CLI
func Execute(initAdditionalCommands func(*cobra.Command)) {
	rootCmd := RootCmd() // Root command from cobra.go

	// Allow additional commands to be added
	if initAdditionalCommands != nil {
		initAdditionalCommands(rootCmd)
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
