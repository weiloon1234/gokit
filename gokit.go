package gokit

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weiloon1234/gokit/database"
	"github.com/weiloon1234/gokit/environment"
	"github.com/weiloon1234/gokit/localization"
	"github.com/weiloon1234/gokit/middleware"
	"github.com/weiloon1234/gokit/storage"
	"io"
	"net/http"
	"os"
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
	if c.TelegramMessageThrottleMinute == 0 {
		return 10 * time.Minute
	}
	return c.TelegramMessageThrottleMinute
}

// Logger instance
var (
	log             = logrus.New()
	sharedConfig    Config
	sharedGinConfig GinConfig
	logCache        = make(map[string]time.Time)
	logMutex        sync.Mutex
)

func Init(config Config) {
	sharedConfig = config

	// Configure logger
	setupLogger(config.Features)

	log.Info("Initializing gokit...")

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

	// Initialize Storage
	if config.StorageProvider != "" {
		err := storage.Init(config.StorageProvider, config.StorageConfig, &config.UploadConfig)
		if err != nil {
			log.Fatalf("Failed to initialize storage: %v", err)
		}
		log.Info("Storage initialized successfully")
	}

	// Initialize Localization
	if config.Features.EnableLocale {
		localization.Init(config.LocalizationConfig)
	}

	log.Info("gokit initialized successfully")
}

func setupLogger(features Features) {
	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Format logs
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Output logs to file and stdout
	if sharedGinConfig.LogFile != "" {
		logFile, err := os.OpenFile(sharedGinConfig.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	} else {
		log.SetOutput(os.Stdout)
	}

	// Add Telegram hook for fatal errors
	if sharedGinConfig.FatalErrorInformTelegram {
		log.AddHook(&TelegramHook{
			Token:    sharedGinConfig.TelegramBotToken,
			ChatID:   sharedGinConfig.TelegramChatId,
			Throttle: sharedGinConfig.GetTelegramMessageThrottleMinute(),
		})
	}
}

func mergeFeatures(custom Features) Features {
	defaults := NewDefaultFeatures()
	if custom.EnableDB {
		defaults.EnableDB = custom.EnableDB
	}
	if !custom.EnableRedis {
		defaults.EnableRedis = custom.EnableRedis
	}
	if custom.EnableLocale {
		defaults.EnableLocale = custom.EnableLocale
	}
	return defaults
}

func NewDefaultFeatures() Features {
	return Features{
		EnableDB:     true,
		EnableRedis:  true,
		EnableLocale: true,
	}
}

type TelegramHook struct {
	Token    string
	ChatID   string
	Throttle time.Duration
}

func (hook *TelegramHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.FatalLevel, logrus.PanicLevel}
}

func (hook *TelegramHook) Fire(entry *logrus.Entry) error {
	message, err := entry.String()
	if err != nil {
		return fmt.Errorf("failed to stringify log entry: %v", err)
	}

	// Throttle messages
	if !throttleLog(message, hook.Throttle) {
		return nil
	}

	return sendTelegramMessage(hook.Token, hook.ChatID, message)
}

func sendTelegramMessage(token, chatID, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	body := fmt.Sprintf(`{"chat_id": "%s", "text": "%s"}`, chatID, message)

	resp, err := http.Post(url, "application/json", bytes.NewBufferString(body))
	if err != nil {
		log.Errorf("Failed to send Telegram message: %v", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Telegram API error: %s", resp.Status)
		return fmt.Errorf("telegram API error: %s", resp.Status)
	}

	return nil
}

func throttleLog(message string, throttle time.Duration) bool {
	logMutex.Lock()
	defer logMutex.Unlock()

	if lastSent, exists := logCache[message]; exists && time.Since(lastSent) < throttle {
		return false
	}

	logCache[message] = time.Now()
	return true
}

func InitRouter(config GinConfig) *gin.Engine {
	sharedGinConfig = config

	// Use Logrus for Gin logs
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(localization.Middleware())
	router.Use(middleware.ErrorLoggingMiddleware())

	return router
}

func GetSharedConfig() Config {
	return sharedConfig
}
