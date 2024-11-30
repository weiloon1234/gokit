package storage

import (
	"fmt"
	"io"
	"mime/multipart"
)

// Storage is the generic interface for file storage
type Storage interface {
	Upload(fileName string, fileContent io.Reader, contentType string) (string, error) // Uploads a file and returns its URL
	Get(fileName string) (string, error)                                               // Gets the file URL
	Delete(fileName string) error                                                      // Deletes the file
}

// UploadConfig holds validation settings for file uploads
type UploadConfig struct {
	AllowedFileTypes  []string // e.g., ["image/png", "image/jpeg", "application/pdf"]
	AllowedImageTypes []string // e.g., ["image/png", "image/jpeg"]
	MaxFileSize       int64    // Max file size in bytes
}

// DefaultUploadConfig provides default settings
var DefaultUploadConfig = UploadConfig{
	AllowedFileTypes:  []string{"image/png", "image/jpeg", "application/pdf"},
	AllowedImageTypes: []string{"image/png", "image/jpeg"},
	MaxFileSize:       5 * 1024 * 1024, // 5 MB
}

// manager holds the initialized storage instance
var manager Storage

// uploadConfig holds the global upload configuration
var uploadConfig UploadConfig

// Init initializes the storage provider and upload configuration globally
func Init(provider string, storageConfig map[string]string, customUploadConfig *UploadConfig) error {
	if provider == "" {
		return fmt.Errorf("storage provider is not configured")
	}

	// Initialize the storage provider
	storage, err := NewStorageProvider(provider, storageConfig)
	if err != nil {
		return err
	}

	manager = storage

	// Initialize upload configuration
	InitUploadConfig(customUploadConfig)

	return nil
}

// InitUploadConfig initializes the global upload configuration
func InitUploadConfig(customConfig *UploadConfig) {
	uploadConfig = DefaultUploadConfig // Start with default settings

	// Merge custom configuration
	if customConfig != nil {
		if len(customConfig.AllowedFileTypes) > 0 {
			uploadConfig.AllowedFileTypes = customConfig.AllowedFileTypes
		}
		if len(customConfig.AllowedImageTypes) > 0 {
			uploadConfig.AllowedImageTypes = customConfig.AllowedImageTypes
		}
		if customConfig.MaxFileSize > 0 {
			uploadConfig.MaxFileSize = customConfig.MaxFileSize
		}
	}
}

// GetManager retrieves the global storage instance
func GetManager() Storage {
	if manager == nil {
		panic("storage manager is not initialized")
	}
	return manager
}

// NewStorageProvider initializes the appropriate storage provider
func NewStorageProvider(provider string, config map[string]string) (Storage, error) {
	switch provider {
	case "s3":
		return NewS3Storage(
			config["bucket"],
			config["region"],
			config["base_url"],
			config["access_key"],
			config["secret_key"],
		), nil
	default:
		return nil, fmt.Errorf("unsupported storage provider: %s", provider)
	}
}

// GetUploadConfig retrieves the global upload configuration
func GetUploadConfig() UploadConfig {
	return uploadConfig
}

// ValidateFile validates the file type and size
func ValidateFile(fileHeader *multipart.FileHeader, config UploadConfig) error {
	if fileHeader.Size > config.MaxFileSize {
		return fmt.Errorf("file size exceeds the maximum limit of %d bytes", config.MaxFileSize)
	}

	fileType := fileHeader.Header.Get("Content-Type")
	if len(config.AllowedFileTypes) > 0 && !contains(config.AllowedFileTypes, fileType) {
		return fmt.Errorf("file type '%s' is not allowed", fileType)
	}

	return nil
}

// ValidateImage validates that the file is an image
func ValidateImage(fileHeader *multipart.FileHeader, config UploadConfig) error {
	fileType := fileHeader.Header.Get("Content-Type")
	if len(config.AllowedImageTypes) > 0 && !contains(config.AllowedImageTypes, fileType) {
		return fmt.Errorf("file type '%s' is not a valid image", fileType)
	}
	return ValidateFile(fileHeader, config)
}

// UploadFile validates and uploads a file
func UploadFile(fileHeader *multipart.FileHeader, file multipart.File) (string, error) {
	// Get the shared configuration
	config := GetUploadConfig()

	// Validate the file
	if err := ValidateFile(fileHeader, config); err != nil {
		return "", err
	}

	// Upload the file
	return manager.Upload(fileHeader.Filename, file, fileHeader.Header.Get("Content-Type"))
}

// UploadImage validates and uploads an image file
func UploadImage(fileHeader *multipart.FileHeader, file multipart.File) (string, error) {
	// Get the shared configuration
	config := GetUploadConfig()

	// Validate the file as an image
	if err := ValidateImage(fileHeader, config); err != nil {
		return "", err
	}

	// Upload the image
	return manager.Upload(fileHeader.Filename, file, fileHeader.Header.Get("Content-Type"))
}

// Helper to check if a slice contains a value
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
