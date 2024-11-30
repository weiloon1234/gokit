package storage

import (
	"path/filepath"
	"strings"
)

// GenerateUniqueFileName appends a unique identifier to the file name
func GenerateUniqueFileName(fileName string, uniqueID string) string {
	ext := filepath.Ext(fileName)
	base := strings.TrimSuffix(fileName, ext)
	return base + "_" + uniqueID + ext
}
