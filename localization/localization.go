package localization

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

type LocaleConfig struct {
	DefaultLanguage    string
	SupportedLanguages []string
	TranslationPaths   []string // Paths to search for locale files
}

var (
	LocaleContextKey   string = "locale"
	defaultLanguage    string
	supportedLanguages []string
	translations       = make(map[string]map[string]string) // locale -> key -> translation
)

func Init(config LocaleConfig) {
	defaultLanguage = config.DefaultLanguage
	supportedLanguages = config.SupportedLanguages

	// Load predefined translations from localization/locales directory
	predefinedPath := "./localization/locales"
	if err := loadPredefinedTranslations(predefinedPath, supportedLanguages); err != nil {
		fmt.Printf("Warning: Failed to load predefined translations: %v\n", err)
	}

	// Load additional translations from custom paths
	for _, path := range config.TranslationPaths {
		if err := loadTranslationsFromPath(path); err != nil {
			fmt.Printf("Warning: Failed to load additional translations from '%s': %v\n", path, err)
		}
	}
}

// loadPredefinedTranslations loads predefined translations from a directory
func loadPredefinedTranslations(path string, languages []string) error {
	for _, lang := range languages {
		filePath := filepath.Join(path, fmt.Sprintf("%s.json", lang))

		// Attempt to load the file
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Warning: Locale file not found for '%s': %s\n", lang, filePath)
			continue // Skip missing files
		}

		if err := loadTranslationsFromFile(filePath, lang); err != nil {
			return fmt.Errorf("failed to load predefined translations for '%s': %v", lang, err)
		}
	}

	// Ensure default language is loaded as fallback
	defaultPath := filepath.Join(path, fmt.Sprintf("%s.json", defaultLanguage))
	if _, exists := translations[defaultLanguage]; !exists {
		if err := loadTranslationsFromFile(defaultPath, defaultLanguage); err != nil {
			return fmt.Errorf("failed to load fallback default language file '%s': %v", defaultPath, err)
		}
	}

	return nil
}

// loadTranslationsFromPath searches for JSON files in a given path and merges them
func loadTranslationsFromPath(path string) error {
	return filepath.WalkDir(path, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(filePath) != ".json" {
			return nil
		}

		locale := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
		return loadTranslationsFromFile(filePath, locale)
	})
}

// loadTranslationsFromFile loads a single locale file and merges its translations
func loadTranslationsFromFile(filePath, locale string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read localization file %s: %v", filePath, err)
	}

	var fileTranslations map[string]string
	if err := json.Unmarshal(data, &fileTranslations); err != nil {
		return fmt.Errorf("failed to parse localization file %s: %v", filePath, err)
	}

	if _, exists := translations[locale]; !exists {
		translations[locale] = make(map[string]string)
	}
	for key, value := range fileTranslations {
		translations[locale][key] = value
	}

	return nil
}

// Middleware adds localization support to Gin HTTP handlers
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if !isSupported(lang) {
			lang = defaultLanguage
		}
		// Store the locale in the Gin context
		c.Set(LocaleContextKey, lang)
		c.Next()
	}
}

// __ translates a key with optional attributes.
func Translate(c *gin.Context, key string, attributes ...map[string]string) string {
	// Check if attributes are provided
	var attr map[string]string
	if len(attributes) > 0 {
		attr = attributes[0] // Use the first attributes map if provided
	}

	// Detect the locale from the Gin context
	locale, exists := c.Get(LocaleContextKey)
	if !exists || locale == "" {
		locale = defaultLanguage // Fallback to default language
	}

	// Perform the translation
	return TranslateKey(locale.(string), key, attr)
}

// TranslateKey a given key into the specified locale
func TranslateKey(locale, key string, attributes map[string]string) string {
	translation := translations[locale][key]
	if translation == "" {
		translation = translations[defaultLanguage][key] // Fallback to default language
	}

	if translation == "" {
		return key // Return the key itself if no translation is found
	}

	// Replace attributes in the translation
	for attrKey, attrValue := range attributes {
		translation = replaceAttribute(translation, attrKey, attrValue)
	}

	return translation
}

// isSupported checks if the language is supported
func isSupported(lang string) bool {
	for _, l := range supportedLanguages {
		if strings.HasPrefix(lang, l) {
			return true
		}
	}
	return false
}

// replaceAttribute replaces :attribute placeholders in the translation
func replaceAttribute(translation, attribute, value string) string {
	placeholder := fmt.Sprintf(":%s", attribute)
	return strings.ReplaceAll(translation, placeholder, value)
}
