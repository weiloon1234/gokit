package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/weiloon1234/gokit/config"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Global Logrus instance
var log = logrus.New()

var (
	logCache = make(map[string]time.Time)
	logMutex sync.Mutex
)

type TelegramHook struct {
	Token    string
	ChatID   string
	Throttle time.Duration
}

func Init(config config.LogConfig) {
	// Set log format
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set output to file and stdout
	if config.LogFile != "" {
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(file)
	} else {
		log.SetOutput(os.Stdout)
	}

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Output logs to file and stdout
	if config.LogFile != "" {
		logFile, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	} else {
		log.SetOutput(os.Stdout)
	}

	if config.FatalErrorInformTelegram {
		log.AddHook(&TelegramHook{
			Token:    config.TelegramBotToken,
			ChatID:   config.TelegramChatId,
			Throttle: config.TelegramMessageThrottleMinute,
		})
	}
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

// GetLogger returns the global logger instance
func GetLogger() *logrus.Logger {
	return log
}
