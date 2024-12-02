package storage

import (
	"fmt"
	"github.com/weiloon1234/gokit/config"
	"io"
	"mime/multipart"
)

// Storage is the generic interface for file storage
type Storage interface {
	Upload(fileName string, fileContent io.Reader, contentType string) (string, error) // Uploads a file and returns its URL
	Get(fileName string) (string, error)                                               // Gets the file URL
	Delete(fileName string) error                                                      // Deletes the file
}

// manager holds the initialized storage instance
var manager Storage

// uploadConfig holds the global upload configuration
var uploadConfig *config.UploadConfig

// Init initializes the storage provider and upload configuration globally
func Init(storageCfg *config.StorageConfig, uploadCfg *config.UploadConfig) error {
	if storageCfg.Provider == "" {
		return fmt.Errorf("storage provider is not configured")
	}

	uploadConfig = uploadCfg

	// Initialize the storage provider
	storage, err := NewStorageProvider(storageCfg)
	if err != nil {
		return err
	}

	manager = storage

	return nil
}

// GetManager retrieves the global storage instance
func GetManager() Storage {
	if manager == nil {
		panic("storage manager is not initialized")
	}
	return manager
}

// NewStorageProvider initializes the appropriate storage provider
func NewStorageProvider(storageCfg *config.StorageConfig) (Storage, error) {
	switch storageCfg.Provider {
	case "s3":
		return NewS3Storage(
			storageCfg.Bucket,
			storageCfg.Region,
			storageCfg.BaseUrl,
			storageCfg.AccessKey,
			storageCfg.SecretKey,
		), nil
	default:
		return nil, fmt.Errorf("unsupported storage provider: %s", storageCfg.Provider)
	}
}

// GetUploadConfig retrieves the global upload configuration
func GetUploadConfig() *config.UploadConfig {
	return uploadConfig
}

// ValidateFile validates the file type and size
func ValidateFile(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > uploadConfig.MaxFileSize {
		return fmt.Errorf("file size exceeds the maximum limit of %d bytes", uploadConfig.MaxFileSize)
	}

	fileType := fileHeader.Header.Get("Content-Type")
	if len(uploadConfig.AllowedFileTypes) > 0 && !contains(uploadConfig.AllowedFileTypes, fileType) {
		return fmt.Errorf("file type '%s' is not allowed", fileType)
	}

	return nil
}

// ValidateImage validates that the file is an image
func ValidateImage(fileHeader *multipart.FileHeader) error {
	fileType := fileHeader.Header.Get("Content-Type")
	if len(uploadConfig.AllowedImageTypes) > 0 && !contains(uploadConfig.AllowedImageTypes, fileType) {
		return fmt.Errorf("file type '%s' is not a valid image", fileType)
	}
	return ValidateFile(fileHeader)
}

// UploadFile validates and uploads a file
func UploadFile(fileHeader *multipart.FileHeader, file multipart.File) (string, error) {
	// Validate the file
	if err := ValidateFile(fileHeader); err != nil {
		return "", err
	}

	// Upload the file
	return manager.Upload(fileHeader.Filename, file, fileHeader.Header.Get("Content-Type"))
}

// UploadImage validates and uploads an image file
func UploadImage(fileHeader *multipart.FileHeader, file multipart.File) (string, error) {
	// Validate the file as an image
	if err := ValidateImage(fileHeader); err != nil {
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
