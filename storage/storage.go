package storage

import (
	"fmt"
	"github.com/weiloon1234/gokit/config"
	"io"
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
