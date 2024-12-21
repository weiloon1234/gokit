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

	goKitSchemaDir := filepath.Join(baseDir, "ent", "schema")
	projectSchemaDir := filepath.Join(currentDir, "ent", "schema")

	entities, err := listEntities(goKitSchemaDir)
	if err != nil {
		fmt.Printf("Error listing base schemas: %v\n", err)
		return
	}

	fmt.Println("Available entities to copy:")
	for _, entity := range entities {
		fmt.Printf(" - %s\n", entity)
	}

	fmt.Print("Enter the entities to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedEntities := strings.Split(input, ",")

	for _, entity := range selectedEntities {
		entity = strings.TrimSpace(entity)
		if entity == "" {
			continue
		}

		baseFile := filepath.Join(goKitSchemaDir, entity+".go")
		targetFile := filepath.Join(projectSchemaDir, entity+".go")

		// Check if the entity already exists in the project
		if utils.FileExists(targetFile) {
			fmt.Printf("Make you you have back up your file if choose to overwrite\n")
			fmt.Printf("Entity %s already exists in the project. Do you want to overwrite it? (y/n): ", entity)
			var response string
			fmt.Scanln(&response)
			response = strings.ToLower(strings.TrimSpace(response))
			if response != "y" && response != "yes" {
				fmt.Printf("Skipping %s.\n", entity)
				continue
			}
		}

		fromModuleName, ok1 := utils.GetModuleName(baseDir)
		toModuleName, ok2 := utils.GetModuleName(currentDir)

		if ok1 != nil || ok2 != nil {
			fmt.Printf("Error getting module name: %v\n", err)
			continue
		}

		baseContent, err := os.ReadFile(baseFile)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", entity, err)
			continue
		}

		updatedContent := []byte(strings.ReplaceAll(string(baseContent), fromModuleName, toModuleName))

		if err := os.WriteFile(targetFile, updatedContent, 0644); err != nil {
			fmt.Printf("Error writing %s: %v\n", entity, err)
			continue
		}
	}

	fmt.Printf("Remember to rebuild entity after copying mixins.\n")
	fmt.Println("Copy Base Entity completed successfully")
}
