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

var CopyEntityHookCmd = &cobra.Command{
	Use:   "entity:copy-hook",
	Short: "Copy entity hook to your project",
	Run:   runCopyEntityHook,
}

func listHooks(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var hooks []string
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") || strings.HasPrefix(file.Name(), "hook.go") {
			continue
		}
		hook := strings.TrimSuffix(file.Name(), ".go")
		hooks = append(hooks, hook)
	}
	return hooks, nil
}

func runCopyEntityHook(cmd *cobra.Command, args []string) {
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

	goKitDir := filepath.Join(goKitRootPath, "ent", "hook")
	projectDir := filepath.Join(projectRootPath, "ent", "hook")

	items, err := listHooks(goKitDir)
	if err != nil {
		fmt.Printf("Error listing hooks: %v\n", err)
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

	selectedItems, err := utils.SelectItems(items, "Select the hooks to copy:")
	if err != nil {
		fmt.Printf("Error selecting seeds: %v\n", err)
		return
	}

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
			fmt.Printf("Hook %s already exists in the project. Do you want to overwrite it? (y/n): ", item)
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

	fmt.Printf("Below try to auto register hook\n")

	//Below for register hooks automatically
	stringToSearchAndAppend := "/** FOR GOKIT AUTO REGISTER ENTITY HOOKS HERE, DON'T EDIT THIS LINE **/"
	missing := false

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
			registerLine := fmt.Sprintf("entClient.Use(hooks.%s())\n", funcName)
			manualAppendMessage := fmt.Sprintf("Please register %s manually.\nRegister in main.go like this\n%s", item, registerLine)

			for _, mainFilePath := range []string{
				filepath.Join(projectRootPath, "cmd", "cli", "main.go"),
				filepath.Join(projectRootPath, "cmd", "web", "main.go"),
			} {
				fileContent, err := os.ReadFile(mainFilePath)
				if err != nil {
					failureMessages = append(failureMessages, fmt.Sprintf("Auto-register: Error reading %s: %v\n%s", mainFilePath, err, manualAppendMessage))
					continue
				}

				originalContent := string(fileContent)

				// Check and add imports if necessary
				importsToAdd := map[string]string{
					fmt.Sprintf("%s/ent/hooks", toModuleName): "",
				}

				originalContent, err = utils.FileAddImports(originalContent, importsToAdd)
				if err != nil {
					failureMessages = append(failureMessages, fmt.Sprintf("Auto-register: Error reading %s: %v\n%s", mainFilePath, err, manualAppendMessage))
					continue
				}

				if strings.Contains(originalContent, registerLine) {
					continue
				} else {
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

						updatedContent := strings.Replace(originalContent, stringToSearchAndAppend, stringToSearchAndAppend+"\n"+indent+registerLine, 1)
						if err := os.WriteFile(mainFilePath, []byte(updatedContent), 0644); err != nil {
							failureMessages = append(failureMessages, fmt.Sprintf("Auto-register: Error writing to %s: %v\n%s", mainFilePath, err, manualAppendMessage))
							continue
						}
						successMessages = append(successMessages, fmt.Sprintf("Auto-register: %s registered in %s", item, mainFilePath))
					} else {
						failureMessages = append(failureMessages, fmt.Sprintf("Auto-register: line not found in %s\n%s", mainFilePath, manualAppendMessage))
						missing = true
						continue
					}
				}
			}
		}

		if len(successMessages) > 0 {
			fmt.Printf("%d hook(s) automatically registered successfully:\n", len(successMessages))
			for _, msg := range successMessages {
				fmt.Printf(" - %s\n", msg)
			}
		}

		if len(failureMessages) > 0 {
			fmt.Printf("%d hook(s) failed to register:\n", len(failureMessages))
			for _, msg := range failureMessages {
				fmt.Printf(" - %s\n", msg)
			}
		}

		fmt.Printf("======================\n")
		if missing {
			fmt.Printf("You main.go missing the placeholder, please add back the placeholder\n")
			fmt.Printf("%s", stringToSearchAndAppend+"\n")
		}
		fmt.Printf("If undefined hook please import your project ent/hook package\n")
		fmt.Printf("======================\n")
	} else {
		fmt.Printf("======================\n")
		fmt.Printf("No hooks copied.\n")
		fmt.Printf("======================\n")
	}

	fmt.Println("Copy Hooks completed successfully")
}
