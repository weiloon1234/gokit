package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/localization"
	"github.com/weiloon1234/gokit/storage"
	"mime/multipart"
)

// ValidateFile validates the file type and size
func ValidateFile(c *gin.Context, fileHeader *multipart.FileHeader, file multipart.File) error {
	cfg := storage.GetUploadConfig()

	if fileHeader.Size > cfg.MaxFileSize {
		return fmt.Errorf(localization.Translate(c, "File size exceeds the maximum limit of :size", map[string]string{
			"size": HumanReadableSize(cfg.MaxFileSize),
		}))
	}

	fileType := fileHeader.Header.Get("Content-Type")
	if len(cfg.AllowedFileTypes) > 0 && !Contains(cfg.AllowedFileTypes, fileType) {
		return fmt.Errorf(localization.Translate(c, "File type :exts is not a valid file or disallowed", map[string]string{
			"exts": ExtractExtension(fileType),
		}))
	}

	return nil
}

// ValidateImage validates that the file is an image
func ValidateImage(c *gin.Context, fileHeader *multipart.FileHeader, file multipart.File) error {
	cfg := storage.GetUploadConfig()

	fileType := fileHeader.Header.Get("Content-Type")
	if len(cfg.AllowedImageTypes) > 0 && !Contains(cfg.AllowedImageTypes, fileType) {
		return fmt.Errorf(localization.Translate(c, "File type :exts is not a valid image or disallowed", map[string]string{
			"exts": ExtractExtension(fileType),
		}))
	}

	return ValidateFile(c, fileHeader, file)
}

// UploadFile validates and uploads a file
func UploadFile(c *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if err := ValidateFile(c, fileHeader, file); err != nil {
		return "", err
	}

	return storage.GetManager().Upload(fileHeader.Filename, file, fileHeader.Header.Get("Content-Type"))
}

// UploadImage validates and uploads an image file
func UploadImage(c *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Validate the file as an image
	if err := ValidateImage(c, fileHeader, file); err != nil {
		return "", err
	}

	// Upload the image
	return storage.GetManager().Upload(fileHeader.Filename, file, fileHeader.Header.Get("Content-Type"))
}

func GetFile(filePath string) string {
	fileUrl, err := storage.GetManager().Get(filePath)
	if err != nil {
		return ""
	}

	return fileUrl
}
