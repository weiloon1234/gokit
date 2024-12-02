package config

import (
	"fmt"
)

type StorageConfig struct {
	Provider  string
	Bucket    string
	Region    string
	BaseUrl   string
	AccessKey string
	SecretKey string
}

func (sc *StorageConfig) BuildConfig() {
	if sc.Provider == "" {
		sc.Provider = "s3"
	}

	if sc.Bucket != "" && sc.BaseUrl != "" && (sc.AccessKey == "" || sc.SecretKey == "") {
		panic(fmt.Sprintf("Access key or secret key for Bucket %q not set", sc.Bucket))
	}
}
