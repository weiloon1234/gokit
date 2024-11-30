package environment

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	AppName  string
	AppEnv   string
	Debug    bool
	LogLevel string
	AppPort  string
	Timezone *time.Location
)

// Init initializes the environment variables and sets up defaults.
func Init() {
	// Load environment variables
	AppName = getEnv("APP_NAME", "MyApp")
	AppEnv = getEnv("APP_ENV", "local")
	Debug = getEnvAsBool("DEBUG", true)
	LogLevel = getEnv("LOG_LEVEL", "debug")
	AppPort = getEnv("PORT", "8080")

	// Set the timezone
	SetTimezone(getEnv("TIMEZONE", "UTC"))

	// Log the loaded configuration (omit sensitive data in production)
	if Debug {
		log.Printf("Environment Loaded: AppName=%s, AppEnv=%s, Debug=%v, LogLevel=%s, Port=%s, Timezone=%s",
			AppName, AppEnv, Debug, LogLevel, AppPort, Timezone.String())
	}
}

// SetTimezone configures the application's timezone based on a string.
func SetTimezone(timezoneStr string) {
	loc, err := time.LoadLocation(timezoneStr)
	if err != nil {
		log.Fatalf("Invalid TIMEZONE value: %s, error: %v", timezoneStr, err)
	}
	Timezone = loc
	time.Local = loc // Set the application's local timezone
	log.Printf("Timezone set to: %s", timezoneStr)
}

// Helper to read an environment variable or fallback to a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper to read an environment variable as a boolean or fallback to a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.ParseBool(valueStr)
		if err != nil {
			log.Printf("Invalid value for %s: %s, defaulting to %v", key, valueStr, defaultValue)
			return defaultValue
		}
		return value
	}
	return defaultValue
}
