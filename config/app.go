package config

type AppConfig struct {
	AppName  string
	AppEnv   string
	AppPort  string
	Debug    bool
	LogLevel string
	Timezone string // Use a string to represent the timezone name
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

	// Set a default timezone if none is provided
	if ac.Timezone == "" {
		ac.Timezone = "UTC"
	}
}
