package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// HumanReadableSize converts a size in bytes to a human-readable string.
func HumanReadableSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d Bytes", size)
	}
}

func GetGoKitRootPath() (string, error) {
	// Obtain the file path of the current file
	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to retrieve caller information")
	}

	// Check if the current directory already contains the go.mod file
	goModPath := filepath.Join(currentFilePath, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		// go.mod file found; return the current directory
		return currentFilePath, nil
	}

	// Start from the directory of the current file
	dir := filepath.Dir(currentFilePath)

	// Traverse upwards to find the go.mod file
	for {
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			// go.mod file found; return the directory
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Reached the root directory without finding go.mod
			break
		}
		dir = parentDir
	}
	return "", errors.New("go.mod file not found in any parent directory")
}

func GetProjectRootPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %v", err)
	}

	// Check if the current directory already contains the go.mod file
	goModPath := filepath.Join(currentDir, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		// go.mod file found; return the current directory
		return currentDir, nil
	}

	// Start from the directory of the current file
	dir := filepath.Dir(currentDir)

	// Traverse upwards to find the go.mod file
	for {
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			// go.mod file found; return the directory
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Reached the root directory without finding go.mod
			break
		}
		dir = parentDir
	}
	return "", errors.New("go.mod file not found in any parent directory")
}

func GetModuleName(rootPath string) (string, error) {
	goFile := filepath.Join(rootPath, "go.mod")

	file, err := os.Open(goFile)
	if err != nil {
		return "", fmt.Errorf("error opening go.mod: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading go.mod: %v\n", err)
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// FileAddImport adds an import to the provided Go source code if it doesn't already exist.
// It parses the source code, adds the import, sorts the imports, and returns the modified code.
func FileAddImports(fileContent string, importsToAdd map[string]string) (string, error) {
	// Parse the source code into an AST
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", fileContent, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// Add imports from importsToAdd map
	for importPath, alias := range importsToAdd {
		if !importExists(file, importPath) {
			addImport(file, importPath, alias)
		}
	}

	// Sort the imports
	ast.SortImports(fset, file)

	// Convert the modified AST back to source code
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, file); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// importExists checks if the specified import path already exists in the AST.
func importExists(f *ast.File, importPath string) bool {
	for _, imp := range f.Imports {
		if strings.Trim(imp.Path.Value, `"`) == importPath {
			return true
		}
	}
	return false
}

// addImport adds a new import to the AST.
func addImport(f *ast.File, importPath, importName string) {
	newImport := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"` + importPath + `"`,
		},
	}
	if importName != "" {
		newImport.Name = ast.NewIdent(importName)
	}

	// Find the import declaration section
	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.IMPORT {
			genDecl.Specs = append(genDecl.Specs, newImport)
			return
		}
	}

	// If no import declaration exists, create a new one
	f.Decls = append([]ast.Decl{&ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{newImport},
	}}, f.Decls...)
}

func FileStringExists(f *ast.File, commentText string) bool {
	for _, cg := range f.Comments {
		for _, c := range cg.List {
			if c.Text == commentText {
				return true
			}
		}
	}
	return false
}
