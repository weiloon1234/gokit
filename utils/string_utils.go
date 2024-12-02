package utils

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

// ExtractExtensions converts a slice of MIME types into a comma-separated list of extensions.
func ExtractExtension(mimeType string) string {
	parts := strings.Split(mimeType, "/")
	var part string
	if len(parts) > 1 {
		part = parts[1]
	}

	return part
}

// ExtractExtensions converts a slice of MIME types into a comma-separated list of extensions.
func ExtractExtensions(mimeTypes []string) string {
	var extensions []string
	for _, mimeType := range mimeTypes {
		// Split the MIME type and take the part after the slash
		parts := strings.Split(mimeType, "/")
		if len(parts) > 1 {
			extensions = append(extensions, parts[1])
		}
	}
	return strings.Join(extensions, ", ")
}
