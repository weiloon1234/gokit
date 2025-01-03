package commands

import (
	"github.com/spf13/cobra"
	database "github.com/weiloon1234/gokit/cli/commands/database"
	entity "github.com/weiloon1234/gokit/cli/commands/entity"
)

// InitCommands adds all core commands to the root command
func InitCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(MigrateCmd)
	rootCmd.AddCommand(SeederCmd)

	//Entity relevant
	rootCmd.AddCommand(entity.CopyEntityCmd)
	rootCmd.AddCommand(entity.CopyEntityMixinCmd)
	rootCmd.AddCommand(entity.CopyEntityHookCmd)

	//Database relevant
	rootCmd.AddCommand(database.CopyDatabaseSeederCmd)
}
