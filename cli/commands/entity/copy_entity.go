package commands

import (
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit/utils"
)

var CopyEntityCmd = &cobra.Command{
	Use:   "entity:copy-entity",
	Short: "Copy base entity to your project",
	Run:   runCopyBaseEntity,
}

func listEntities(dir string) ([]string, error) {
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

func runCopyBaseEntity(cmd *cobra.Command, args []string) {
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

	goKitDir := filepath.Join(goKitRootPath, "ent", "schema")
	projectDir := filepath.Join(projectRootPath, "ent", "schema")

	items, err := listEntities(goKitDir)
	if err != nil {
		fmt.Printf("Error listing entity: %v\n", err)
		return
	}

	fromModuleName, ok1 := utils.GetModuleName(goKitRootPath)
	toModuleName, ok2 := utils.GetModuleName(projectRootPath)

	if ok1 != nil || ok2 != nil {
		fmt.Printf("Error getting module name: %v\n", err)
		return
	}

	// Ensure the target directory exists
	if err := os.MkdirAll(projectDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", projectDir, err)
	}

	fmt.Println("Available entities to copy:")
	for _, item := range items {
		fmt.Printf(" - %s\n", item)
	}

	fmt.Print("Enter the entities to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedItems := strings.Split(input, ",")

	for _, item := range selectedItems {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		baseFile := filepath.Join(goKitRootPath, item+".go")
		targetFile := filepath.Join(projectRootPath, item+".go")

		// Check if the entity already exists in the project
		if utils.FileExists(targetFile) {
			fmt.Printf("Make you you have back up your file if choose to overwrite\n")
			fmt.Printf("Entity %s already exists in the project. Do you want to overwrite it? (y/n): ", item)
			var response string
			fmt.Scanln(&response)
			response = strings.ToLower(strings.TrimSpace(response))
			if response != "y" && response != "yes" {
				fmt.Printf("Skipping %s.\n", item)
				continue
			}
		}

		baseContent, err := os.ReadFile(baseFile)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", item, err)
			continue
		}

		updatedContent := []byte(strings.ReplaceAll(string(baseContent), fromModuleName, toModuleName))

		if err := os.WriteFile(targetFile, updatedContent, 0644); err != nil {
			fmt.Printf("Error writing %s: %v\n", item, err)
			continue
		}

		fmt.Printf("Successfully copied %s to the project.\n", item)
	}

	fmt.Printf("Remember to rebuild entity after copying mixins.\n")
	fmt.Println("Copy Base Entity completed successfully")
}
