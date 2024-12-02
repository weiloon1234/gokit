package config

import (
	"fmt"
	"time"
)

type LogConfig struct {
	LogFile                       string
	FatalErrorInformTelegram      bool
	TelegramBotToken              string
	TelegramChatId                string
	TelegramMessageThrottleMinute time.Duration
}

func (c *LogConfig) BuildConfig() {
	if c.LogFile == "" {
		c.LogFile = "tmp/server.log"
	}

	if c.FatalErrorInformTelegram != true {
		c.FatalErrorInformTelegram = false
	}

	//Minimum 1 minute
	if c.TelegramMessageThrottleMinute <= 0 {
		c.TelegramMessageThrottleMinute = 1 * time.Minute
	}

	if c.FatalErrorInformTelegram {
		if c.TelegramBotToken == "" || c.TelegramChatId == "" {
			panic(fmt.Sprintf("Incomplete Telegram configuration: TelegramBotToken=%q, TelegramChatId=%q, TelegramMessageThrottleMinute=%v",
				c.TelegramBotToken, c.TelegramChatId, c.TelegramMessageThrottleMinute))
		}
	}
}
