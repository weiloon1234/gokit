package gokit

import (
	"github.com/weiloon1234/gokit/database"
	"github.com/weiloon1234/gokit/environment"
	"github.com/weiloon1234/gokit/localization"
	"github.com/weiloon1234/gokit/storage"
	"log"
)

// Features to toggle optional components
type Features struct {
	EnableDB     bool
	EnableRedis  bool
	EnableLocale bool
}

// DBConfig Re-export
type DBConfig = database.DBConfig

// RedisConfig Re-export
type RedisConfig = database.RedisConfig

// LocaleConfig Re-export
type LocaleConfig = localization.LocaleConfig

// UploadConfig Re-export
type UploadConfig = storage.UploadConfig

// Config holds configuration for all components
type Config struct {
	DBConfig           DBConfig
	RedisConfig        RedisConfig
	LocalizationConfig LocaleConfig
	Timezone           string
	StorageProvider    string
	StorageConfig      map[string]string
	UploadConfig       UploadConfig
	Features           Features
}

// Shared configuration accessible by other packages
var sharedConfig Config

// NewDefaultFeatures returns Features with all flags set to true
func NewDefaultFeatures() Features {
	return Features{
		EnableDB:     true,
		EnableRedis:  true,
		EnableLocale: true,
	}
}

// mergeFeatures merges provided Features with default values
func mergeFeatures(custom Features) Features {
	defaults := NewDefaultFeatures()
	if custom.EnableDB {
		defaults.EnableDB = custom.EnableDB
	}
	if !custom.EnableRedis { // Default is true; keep it false if explicitly disabled
		defaults.EnableRedis = custom.EnableRedis
	}
	if custom.EnableLocale {
		defaults.EnableLocale = custom.EnableLocale
	}
	return defaults
}

// Init initializes all core components
func Init(config Config) {
	// Store the config for global access
	sharedConfig = config

	// Merge custom features with default values
	config.Features = mergeFeatures(config.Features)

	// Initialize environment variables
	environment.Init()

	// Initialize Database
	if config.Features.EnableDB {
		if err := database.Init(config.DBConfig); err != nil {
			log.Fatalf("Failed to initialize DB: %v", err)
		}
	}

	// Initialize Redis
	if config.Features.EnableRedis {
		if err := database.InitRedis(config.RedisConfig); err != nil {
			log.Fatalf("Failed to initialize Redis: %v", err)
		}
	}

	// Initialize Storage (pass relevant configuration)
	if config.StorageProvider != "" {
		err := storage.Init(config.StorageProvider, config.StorageConfig, &config.UploadConfig)
		if err != nil {
			log.Fatalf("Failed to initialize storage: %v", err)
		}
		log.Println("Storage initialized successfully")
	}

	// Initialize Localization
	if config.Features.EnableLocale {
		localization.Init(config.LocalizationConfig)
	}

	log.Println("gokit initialized successfully")
}

// GetSharedConfig exposes the stored configuration for other packages
func GetSharedConfig() Config {
	return sharedConfig
}
