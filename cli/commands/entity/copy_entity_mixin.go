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
	goKitRootPath, err := utils.GetGoKitRootPath()
	if err != nil {
		fmt.Printf("Error getting GoKit root path: %v\n", err)
		return
	}

	projectRootPath, err := utils.GetProjectRootPath()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}

	goKitDir := filepath.Join(goKitRootPath, "ent", "mixin")
	projectDir := filepath.Join(projectRootPath, "ent", "mixin")

	items, err := listMixins(goKitDir)
	if err != nil {
		fmt.Printf("Error listing mixins: %v\n", err)
		return
	}

	// Ensure the target directory exists
	if err := os.MkdirAll(projectDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", projectDir, err)
	}

	selectedItems, err := utils.SelectItems(items, "Select mixins to copy:")
	if err != nil {
		fmt.Printf("Error selecting seeds: %v\n", err)
		return
	}

	for _, item := range selectedItems {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		baseFile := filepath.Join(goKitDir, item+".go")
		targetFile := filepath.Join(projectDir, item+".go")

		if utils.FileExists(targetFile) {
			fmt.Printf("Make you you have back up your file if choose to overwrite\n")
			fmt.Printf("Mixin %s already exists in the project. Do you want to overwrite it? (y/n): ", item)
			var response string
			fmt.Scanln(&response)
			response = strings.ToLower(strings.TrimSpace(response))
			if response != "y" && response != "yes" {
				fmt.Printf("Skipping %s.\n", item)
				continue
			}
		}

		// Copy the schema file
		if err := utils.CopyFile(baseFile, targetFile); err != nil {
			fmt.Printf("Error copying %s: %v\n", item, err)
		} else {
			fmt.Printf("Successfully copied %s to the project.\n", item)
		}
	}

	fmt.Printf("Remember to rebuild entity after copying mixins.\n")
	fmt.Println("Copy Entity Mixin completed successfully")
}
