package config

type UploadConfig struct {
	AllowedFileTypes  []string // e.g., ["image/png", "image/jpeg", "application/pdf"]
	AllowedImageTypes []string // e.g., ["image/png", "image/jpeg"]
	MaxFileSize       int64    // Max file size in bytes
}

func (u *UploadConfig) BuildConfig() {
	if len(u.AllowedFileTypes) == 0 {
		u.AllowedImageTypes = []string{"image/png", "image/jpeg", "application/pdf"}
	}

	if len(u.AllowedImageTypes) == 0 {
		u.AllowedImageTypes = []string{"image/png", "image/jpeg", "image/jpg", "image/gif", "image/bmp", "image/tiff"}
	}

	//How to do this MaxFileSize
	if u.MaxFileSize == 0 {
		u.MaxFileSize = 5 * 1024 * 1024 // 5 MB
	}
}
