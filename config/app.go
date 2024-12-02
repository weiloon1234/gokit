package config

import (
	"time"
)

type AppConfig struct {
	AppName  string
	AppEnv   string
	AppPort  string
	Debug    bool
	LogLevel string
	Timezone time.Location
}

func (ac *AppConfig) BuildConfig() {
	if ac.AppName == "" {
		ac.AppName = "GoKit"
	}

	if ac.AppEnv == "" {
		ac.AppEnv = "development"
	}

	if ac.AppPort == "" {
		ac.AppPort = "8080"
	}

	if ac.Debug {
		ac.Debug = true
	}

	if ac.LogLevel == "" {
		ac.LogLevel = "info"
	}
}
