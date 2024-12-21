package utils

import (
	"bufio"
	"errors"
	"fmt"
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
