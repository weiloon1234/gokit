package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/weiloon1234/gokit/utils"
	"time"
)

type Config struct {
	App                AppConfig
	LogConfig          LogConfig
	StorageConfig      StorageConfig
	UploadConfig       UploadConfig
	Features           Features
	DBConfig           DBConfig
	RedisConfig        RedisConfig
	LocalizationConfig LocaleConfig
}

func (cfg *Config) BuildConfig() {
	cfg.App.BuildConfig()
	cfg.LogConfig.BuildConfig()
	cfg.StorageConfig.BuildConfig()
	cfg.UploadConfig.BuildConfig()
	cfg.Features.BuildConfig()
	if utils.BoolValue(cfg.Features.EnableDB) {
		cfg.DBConfig.BuildConfig()
	}
	if utils.BoolValue(cfg.Features.EnableRedis) {
		cfg.RedisConfig.BuildConfig()
	}
	if utils.BoolValue(cfg.Features.EnableLocale) {
		cfg.LocalizationConfig.BuildConfig()
	}
}

func (cfg *Config) BuildApp() {
	// SetTimezone configures the application's timezone based on a string.
	timezoneStr := cfg.App.Timezone.String()
	loc, err := time.LoadLocation(timezoneStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid TIMEZONE value: %q, error: %v", timezoneStr, err))
	}

	time.Local = loc

	log.Printf("Timezone set to: %s", timezoneStr)
}
