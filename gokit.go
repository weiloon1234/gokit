package gokit

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/database"
	"github.com/weiloon1234/gokit/environment"
	"github.com/weiloon1234/gokit/localization"
	"github.com/weiloon1234/gokit/middleware"
	"github.com/weiloon1234/gokit/storage"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
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
	DBConfig           *DBConfig
	RedisConfig        *RedisConfig
	LocalizationConfig LocaleConfig
	Timezone           string
	StorageProvider    string
	StorageConfig      map[string]string
	UploadConfig       UploadConfig
	Features           Features
}

type GinConfig struct {
	LogFile                       string
	FatalErrorInformTelegram      bool
	TelegramBotToken              string
	TelegramChatId                string
	TelegramMessageThrottleMinute time.Duration
}

func (c GinConfig) GetTelegramMessageThrottleMinute() time.Duration {
	// If not set, return the default value
	if c.TelegramMessageThrottleMinute == 0 {
		return 10 * time.Minute // Default throttle duration
	}
	return c.TelegramMessageThrottleMinute
}

// Shared configuration accessible by other packages
var sharedConfig Config
var sharedGinConfig GinConfig

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

var (
	logCache      = make(map[string]time.Time)
	logCacheMutex sync.Mutex
)

type customLogger struct{}

func (cl customLogger) Write(p []byte) (n int, err error) {
	// Extract file and line for throttling
	_, file, line, ok := runtime.Caller(3)
	logMessage := fmt.Sprintf("%s:%d - %s", file, line, string(p))
	if !ok {
		logMessage = string(p)
	}

	// Log locally
	log.Print(logMessage)

	// Send Telegram message if not throttled
	if throttleLog(logMessage) {
		go sendTelegramMessage(logMessage)
	}
	return len(p), nil
}

// sendTelegramMessage sends a message via Telegram Bot API.
func sendTelegramMessage(message string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", sharedGinConfig.TelegramBotToken)
	body := fmt.Sprintf(`{"chat_id": "%s", "text": "%s"}`, sharedGinConfig.TelegramChatId, message)

	// Send the request
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(body))
	if err != nil {
		log.Printf("Failed to send Telegram message: %v", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close Telegram message: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Telegram API error: %s", resp.Status)
	}
}

// throttleLog ensures the same log message is only sent once every `messageThrottle`.
func throttleLog(message string) bool {
	logCacheMutex.Lock()
	defer logCacheMutex.Unlock()

	// Check if the message exists and is within the throttle time
	if lastSent, exists := logCache[message]; exists && time.Since(lastSent) < sharedGinConfig.GetTelegramMessageThrottleMinute() {
		return false
	}

	// Update the cache
	logCache[message] = time.Now()
	return true
}

func InitRouter(config GinConfig) *gin.Engine {
	sharedGinConfig = config

	// Set up custom logging
	if config.LogFile != "" {
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	}

	// Ensure customLogger implements io.Writer
	gin.DefaultWriter = customLogger{}
	gin.DefaultErrorWriter = customLogger{}

	// Use gin.New() if you want full control over the middlewares
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(localization.Middleware())
	router.Use(middleware.ErrorLoggingMiddleware())

	return router
}

// GetSharedConfig exposes the stored configuration for other packages
func GetSharedConfig() Config {
	return sharedConfig
}
