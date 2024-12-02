package config

import (
	"fmt"
	"github.com/weiloon1234/gokit/utils"
	"time"
)

type LogConfig struct {
	LogFile                       string
	FatalErrorInformTelegram      *bool
	TelegramBotToken              string
	TelegramChatId                string
	TelegramMessageThrottleMinute *time.Duration
}

func (c *LogConfig) BuildConfig() {
	if c.LogFile == "" {
		c.LogFile = "tmp/server.log"
	}

	if c.FatalErrorInformTelegram == nil {
		defaultFatalErrorInformTelegram := true
		c.FatalErrorInformTelegram = &defaultFatalErrorInformTelegram
	}

	if utils.BoolValue(c.FatalErrorInformTelegram) {
		if c.TelegramBotToken == "" || c.TelegramChatId == "" || c.TelegramMessageThrottleMinute == nil {
			panic(fmt.Sprintf("Incomplete Telegram configuration: TelegramBotToken=%q, TelegramChatId=%q, TelegramMessageThrottleMinute=%v",
				c.TelegramBotToken, c.TelegramChatId, c.TelegramMessageThrottleMinute))
		}
	}
}
