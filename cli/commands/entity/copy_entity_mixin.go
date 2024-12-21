package commands

import (
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit/utils"
)

var CopyEntityMixinCmd = &cobra.Command{
	Use:   "entity:copy-mixin",
	Short: "Copy entity mixin to your project",
	Run:   runCopyEntityMixin,
}

func listMixins(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var entities []string
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		entity := strings.TrimSuffix(file.Name(), ".go")
		entities = append(entities, entity)
	}
	return entities, nil
}

func runCopyEntityMixin(cmd *cobra.Command, args []string) {
	baseDir, ok := utils.GetGoKitRootPath()
	if ok != nil {
		fmt.Printf("Error getting GoKit root path: %v\n", ok)
		return
	}

	currentDir, ok2 := utils.GetProjectRootPath()
	if ok2 != nil {
		fmt.Printf("Error getting current working directory: %v\n", ok2)
		return
	}

	goKitMixinDir := filepath.Join(baseDir, "ent", "mixin")
	projectMixinDir := filepath.Join(currentDir, "ent", "mixin")

	mixins, err := listMixins(goKitMixinDir)
	if err != nil {
		fmt.Printf("Error listing base schemas: %v\n", err)
		return
	}

	fmt.Println("Available entities to copy:")
	for _, mixin := range mixins {
		fmt.Printf(" - %s\n", mixin)
	}

	fmt.Print("Enter the mixin to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedMixins := strings.Split(input, ",")

	for _, mixin := range selectedMixins {
		mixin = strings.TrimSpace(mixin)
		if mixin == "" {
			continue
		}

		baseFile := filepath.Join(goKitMixinDir, mixin+".go")
		targetFile := filepath.Join(projectMixinDir, mixin+".go")

		// Check if the mixin already exists in the project
		if utils.FileExists(targetFile) {
			fmt.Printf("Make you you have back up your file if choose to overwrite\n")
			fmt.Printf("Mixin %s already exists in the project. Do you want to overwrite it? (y/n): ", mixin)
			var response string
			fmt.Scanln(&response)
			response = strings.ToLower(strings.TrimSpace(response))
			if response != "y" && response != "yes" {
				fmt.Printf("Skipping %s.\n", mixin)
				continue
			}
		}

		// Copy the schema file
		if err := utils.CopyFile(baseFile, targetFile); err != nil {
			fmt.Printf("Error copying %s: %v\n", mixin, err)
		} else {
			fmt.Printf("Successfully copied %s to the project.\n", mixin)
		}
	}

	fmt.Printf("Remember to rebuild entity after copying mixins.\n")
	fmt.Println("Copy Entity Mixin completed successfully")
}
