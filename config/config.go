package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/weiloon1234/gokit/utils"
	"time"
)

type Config struct {
	AppConfig          AppConfig
	LogConfig          LogConfig
	StorageConfig      StorageConfig
	UploadConfig       UploadConfig
	FeatureConfig      FeatureConfig
	DBConfig           DBConfig
	RedisConfig        RedisConfig
	LocalizationConfig LocaleConfig
}

func (cfg *Config) BuildConfig() {
	cfg.AppConfig.BuildConfig()
	cfg.LogConfig.BuildConfig()
	cfg.StorageConfig.BuildConfig()
	cfg.UploadConfig.BuildConfig()
	cfg.FeatureConfig.BuildConfig()
	if utils.BoolValue(cfg.FeatureConfig.EnableDB) {
		cfg.DBConfig.BuildConfig()
	}
	if utils.BoolValue(cfg.FeatureConfig.EnableRedis) {
		cfg.RedisConfig.BuildConfig()
	}
	if utils.BoolValue(cfg.FeatureConfig.EnableLocale) {
		cfg.LocalizationConfig.BuildConfig()
	}
}

func (cfg *Config) BuildApp() {
	// SetTimezone configures the application's timezone based on a string.
	timezoneStr := cfg.AppConfig.Timezone.String()
	loc, err := time.LoadLocation(timezoneStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid TIMEZONE value: %q, error: %v", timezoneStr, err))
	}

	time.Local = loc

	log.Printf("Timezone set to: %s", timezoneStr)
}
