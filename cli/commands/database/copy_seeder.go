package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var CopyDatabaseSeederCmd = &cobra.Command{
	Use:   "database:copy-seeder",
	Short: "Copy database seeder to your project",
	Run:   runCopyDatabaseSeeder,
}

func listSeeds(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var hooks []string
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		hook := strings.TrimSuffix(file.Name(), ".go")
		hooks = append(hooks, hook)
	}
	return hooks, nil
}

func runCopyDatabaseSeeder(cmd *cobra.Command, args []string) {
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

	goKitDir := filepath.Join(goKitRootPath, "seeds")
	projectDir := filepath.Join(projectRootPath, "seeds")

	items, err := listSeeds(goKitDir)
	if err != nil {
		fmt.Printf("Error listing seeds: %v\n", err)
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

	fmt.Println("Available seeds to copy:")
	for _, item := range items {
		fmt.Printf(" - %s\n", item)
	}

	fmt.Print("Enter the seeds to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedItems := strings.Split(input, ",")

	var successItems []string
	for _, item := range selectedItems {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		baseFile := filepath.Join(goKitDir, item+".go")
		targetFile := filepath.Join(projectDir, item+".go")

		if utils.FileExists(targetFile) {
			fmt.Printf("Make you you have back up your file if choose to overwrite\n")
			fmt.Printf("Seed %s already exists in the project. Do you want to overwrite it? (y/n): ", item)
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
		successItems = append(successItems, item)
	}

	stringToSearchAndAppend := "/** FOR GOKIT AUTO REGISTER SEEDER HERE, DON'T EDIT THIS LINE **/"

	if len(successItems) > 0 {
		fmt.Printf("======================\n")
		var successMessages []string
		var failureMessages []string

		for _, item := range successItems {
			funcName := strings.ReplaceAll(
				cases.Title(language.English).String(strings.ReplaceAll(item, "_", " ")),
				" ",
				"",
			)
			registerLine := fmt.Sprintf("goKitCommand.RegisterSeeder(\"%s\", func() { seeds.%s(entClient) })", item, funcName)
			manualAppendMessage := fmt.Sprintf("Please register %s manually.\nRegister in main.go like this\n%s", item, registerLine)

			for _, mainFilePath := range []string{
				filepath.Join(projectRootPath, "cmd", "cli", "main.go"),
			} {
				fileContent, err := os.ReadFile(mainFilePath)
				if err != nil {
					failureMessages = append(failureMessages, fmt.Sprintf("Error reading %s: %v\n%s", mainFilePath, err, manualAppendMessage))
					continue
				}

				originalContent := string(fileContent)

				// Check and add imports if necessary
				importsToAdd := map[string]string{
					fmt.Sprintf("%s/cli/commands", fromModuleName): "goKitCommand",
					fmt.Sprintf("%s/seeds", toModuleName):          "",
				}

				originalContent, err = utils.FileAddImports(originalContent, importsToAdd)
				if err != nil {
					failureMessages = append(failureMessages, fmt.Sprintf("Error reading %s: %v\n%s", mainFilePath, err, manualAppendMessage))
					continue
				}

				if strings.Contains(originalContent, stringToSearchAndAppend) {
					// Find the indentation before the stringToSearchAndAppend
					lines := strings.Split(originalContent, "\n")
					var indent string
					for _, line := range lines {
						if strings.Contains(line, stringToSearchAndAppend) {
							indent = line[:strings.Index(line, stringToSearchAndAppend)]
							break
						}
					}

					updatedContent := strings.Replace(originalContent, stringToSearchAndAppend, stringToSearchAndAppend+"\n"+indent+registerLine+"\n", 1)
					if err := os.WriteFile(mainFilePath, []byte(updatedContent), 0644); err != nil {
						failureMessages = append(failureMessages, fmt.Sprintf("Error writing to %s: %v\n%s", mainFilePath, err, manualAppendMessage))
						continue
					}
					successMessages = append(successMessages, fmt.Sprintf("%s registered in %s", item, mainFilePath))
				} else {
					failureMessages = append(failureMessages, fmt.Sprintf("Auto-register line not found in %s\n%s", mainFilePath, manualAppendMessage))
					continue
				}
			}
		}

		if len(successMessages) > 0 {
			fmt.Printf("%d seed(s) automatically registered successfully:\n", len(successMessages))
			for _, msg := range successMessages {
				fmt.Printf(" - %s\n", msg)
			}
		}

		if len(failureMessages) > 0 {
			fmt.Printf("%d seed(s) failed to register:\n", len(failureMessages))
			for _, msg := range failureMessages {
				fmt.Printf(" - %s\n", msg)
			}
		}

		fmt.Printf("======================\n")
		fmt.Printf("If undefined goKitCommand please import goKitCommand \"github.com/weiloon1234/gokit/cli/commands\"\n")
		fmt.Printf("If undefined seeds please import your project seeds package\n")
		fmt.Printf("======================\n")
	} else {
		fmt.Printf("======================\n")
		fmt.Printf("No seeds copied.\n")
		fmt.Printf("======================\n")
	}

	fmt.Println("Copy Seeds completed successfully")
}
