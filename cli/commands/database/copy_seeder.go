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
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") || strings.HasPrefix(file.Name(), "hook.go") {
			continue
		}
		hook := strings.TrimSuffix(file.Name(), ".go")
		hooks = append(hooks, hook)
	}
	return hooks, nil
}

func runCopyDatabaseSeeder(cmd *cobra.Command, args []string) {
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

	goKitSeedDir := filepath.Join(baseDir, "seeds")
	projectSeedDir := filepath.Join(currentDir, "seeds")

	seeds, err := listSeeds(goKitSeedDir)
	if err != nil {
		fmt.Printf("Error listing seeds: %v\n", err)
		return
	}

	fmt.Println("Available seeds to copy:")
	for _, seed := range seeds {
		fmt.Printf(" - %s\n", seed)
	}

	fmt.Print("Enter the seeds to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedSeeds := strings.Split(input, ",")

	var successSeeds []string
	for _, seed := range selectedSeeds {
		seed = strings.TrimSpace(seed)
		if seed == "" {
			continue
		}

		baseFile := filepath.Join(goKitSeedDir, seed+".go")
		targetFile := filepath.Join(projectSeedDir, seed+".go")

		if utils.FileExists(targetFile) {
			fmt.Printf("Make you you have back up your file if choose to overwrite\n")
			fmt.Printf("Seed %s already exists in the project. Do you want to overwrite it? (y/n): ", seed)
			var response string
			fmt.Scanln(&response)
			response = strings.ToLower(strings.TrimSpace(response))
			if response != "y" && response != "yes" {
				fmt.Printf("Skipping %s.\n", seed)
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
			fmt.Printf("Error reading %s: %v\n", seed, err)
			continue
		}

		updatedContent := []byte(strings.ReplaceAll(string(baseContent), fromModuleName, toModuleName))

		if err := os.WriteFile(targetFile, updatedContent, 0644); err != nil {
			fmt.Printf("Error writing %s: %v\n", seed, err)
			continue
		}

		fmt.Printf("Successfully copied %s to the project.\n", seed)
		successSeeds = append(successSeeds, seed)
	}

	if len(successSeeds) > 0 {
		fmt.Printf("======================\n")
		var successMessages []string
		var failureMessages []string

		for _, seed := range successSeeds {
			hookFuncName := strings.ReplaceAll(
				cases.Title(language.English).String(strings.ReplaceAll(seed, "_", " ")),
				" ",
				"",
			)
			registerLine := fmt.Sprintf("gokitCommand.RegisterSeeder(\"%s\", seeds.%s)\n", seed, hookFuncName)
			autoRegisterLine := "/** FOR GOKIT AUTO REGISTER SEEDER HERE, DON'T EDIT THIS LINE **/"

			for _, mainFilePath := range []string{
				filepath.Join(currentDir, "cmd", "cli", "main.go"),
			} {
				mainFileContent, err := os.ReadFile(mainFilePath)
				if err != nil {
					failureMessages = append(failureMessages, fmt.Sprintf("Error reading %s: %v\nPlease register %s manually.\nRegister in main.go like this\n%s", mainFilePath, err, seed, registerLine))
					continue
				}

				contentStr := string(mainFileContent)
				if strings.Contains(contentStr, autoRegisterLine) {
					// Find the indentation before the autoRegisterLine
					lines := strings.Split(contentStr, "\n")
					var indent string
					for _, line := range lines {
						if strings.Contains(line, autoRegisterLine) {
							indent = line[:strings.Index(line, autoRegisterLine)]
							break
						}
					}

					updatedContent := strings.Replace(contentStr, autoRegisterLine, autoRegisterLine+"\n"+indent+registerLine, 1)
					if err := os.WriteFile(mainFilePath, []byte(updatedContent), 0644); err != nil {
						failureMessages = append(failureMessages, fmt.Sprintf("Error writing to %s: %v\nPlease register %s manually.\nRegister in main.go like this\n%s", mainFilePath, err, seed, registerLine))
						continue
					}
					successMessages = append(successMessages, fmt.Sprintf("%s registered in %s", seed, mainFilePath))
				} else {
					failureMessages = append(failureMessages, fmt.Sprintf("Auto-register line not found in %s\nPlease register %s manually.\nRegister in main.go like this\n%s", mainFilePath, seed, registerLine))
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
			fmt.Printf("Please register the hooks manually in the respective main.go files.\n")
		}

		fmt.Printf("======================\n")
		fmt.Printf("Remember to rebuild entity after copying hooks.\n")
	} else {
		fmt.Printf("======================\n")
		fmt.Printf("No hooks copied.\n")
		fmt.Printf("======================\n")
	}

	fmt.Println("Copy Seeds completed successfully")
}
