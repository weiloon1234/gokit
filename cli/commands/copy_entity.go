package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit-base-entity/helpers"
)

var CopyEntityCmd = &cobra.Command{
	Use:   "copy-base-entity",
	Short: "Copy base schemas to your project",
	Run:   runCopyBaseEntity,
}

func runCopyBaseEntity(cmd *cobra.Command, args []string) {
	helpers.RunCopyBaseEntity(cmd, args)
	fmt.Println("Copy Base Entity completed successfully")
}
