package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit-base-entity/helpers"
)

var CopyEntityHookCmd = &cobra.Command{
	Use:   "copy-ent-hook",
	Short: "Copy base schemas to your project",
	Run:   runCopyBaseEntityHook,
}

func runCopyBaseEntityHook(cmd *cobra.Command, args []string) {
	helpers.RunCopyBaseEntityHook(cmd, args)
	fmt.Println("Copy Base Entity hook completed successfully")
}
