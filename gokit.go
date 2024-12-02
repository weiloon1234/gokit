package gokit

import (
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/config"
	"github.com/weiloon1234/gokit/database"
	"github.com/weiloon1234/gokit/localization"
	"github.com/weiloon1234/gokit/logger"
	"github.com/weiloon1234/gokit/middleware"
	"github.com/weiloon1234/gokit/storage"
	"github.com/weiloon1234/gokit/utils"
	"log"
)

// Logger instance
var sharedConfig config.Config

func Init(config config.Config) {
	config.BuildConfig()
	sharedConfig = config

	// Configure logger
	logger.Init(config.LogConfig)
	logger.GetLogger().Info("Initializing gokit...")

	// Start App, some setting
	config.BuildApp()

	if utils.BoolValue(config.Features.EnableDB) {
		if err := database.Init(&config.DBConfig); err != nil {
			log.Fatalf("Failed to initialize DB: %v", err)
		}
	}

	// Initialize Redis
	if utils.BoolValue(config.Features.EnableRedis) {
		if err := database.InitRedis(&config.RedisConfig); err != nil {
			log.Fatalf("Failed to initialize Redis: %v", err)
		}
	}

	// Initialize Localization
	if utils.BoolValue(config.Features.EnableLocale) {
		localization.Init(&config.LocalizationConfig)
	}

	// Initialize Storage
	if config.StorageConfig.Provider != "" {
		err := storage.Init(&config.StorageConfig, &config.UploadConfig)
		if err != nil {
			log.Fatalf("Failed to initialize storage: %v", err)
		}
		logger.GetLogger().Info("Storage initialized successfully")
	}

	logger.GetLogger().Info("gokit initialized successfully")
}

func InitRouter(config config.Config) *gin.Engine {
	// Use Logrus for Gin logs
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()

	router := gin.New()
	router.Use(gin.Recovery())
	if utils.BoolValue(config.Features.EnableLocale) {
		router.Use(localization.Middleware())
	}
	router.Use(middleware.RealIPMiddleware())
	router.Use(middleware.ErrorLoggingMiddleware())

	return router
}
