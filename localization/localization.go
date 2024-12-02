package localization

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/config"
	"os"
	"path/filepath"
	"strings"
)

// Embed predefined locale files
//
//go:embed locales/*.json
var embeddedLocales embed.FS

var translations = make(map[string]map[string]string) // locale -> key -> translation
var (
	GlobalLocaleConfig *config.LocaleConfig
	LocaleContextKey   string = "locale"
)

func Init(cfg *config.LocaleConfig) {
	if cfg == nil {
		panic("Localization config cannot be nil")
	}

	// Assign to global variable
	GlobalLocaleConfig = cfg

	// Load predefined translations embedded in the gokit module
	if err := loadEmbeddedTranslations(cfg.SupportedLanguages); err != nil {
		fmt.Printf("Warning: Failed to load embedded translations: %v\n", err)
	}

	// Load additional project-level translations from paths
	for _, path := range cfg.TranslationPaths {
		if err := loadTranslationsFromPath(path); err != nil {
			fmt.Printf("Warning: Failed to load additional translations from '%s': %v\n", path, err)
		}
	}
}

// loadEmbeddedTranslations loads translations embedded in the gokit module
func loadEmbeddedTranslations(languages []string) error {
	for _, lang := range languages {
		filePath := fmt.Sprintf("locales/%s.json", lang) // Path relative to embed

		data, err := embeddedLocales.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Warning: Embedded locale file not found for '%s': %s\n", lang, filePath)
			continue
		}

		if err := mergeTranslations(data, lang); err != nil {
			return fmt.Errorf("failed to load embedded translations for '%s': %v", lang, err)
		}
	}
	return nil
}

// loadTranslationsFromPath loads project-level translations from the provided path
func loadTranslationsFromPath(path string) error {
	return filepath.WalkDir(path, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(filePath) != ".json" {
			return nil
		}

		locale := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read localization file %s: %v", filePath, err)
		}

		return mergeTranslations(data, locale)
	})
}

// mergeTranslations merges a JSON translation file into the translations map
func mergeTranslations(data []byte, locale string) error {
	var fileTranslations map[string]string
	if err := json.Unmarshal(data, &fileTranslations); err != nil {
		return fmt.Errorf("failed to parse localization file for '%s': %v", locale, err)
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
			lang = GlobalLocaleConfig.DefaultLanguage
		}
		// Store the locale in the Gin context
		c.Set(LocaleContextKey, lang)
		c.Next()
	}
}

// Translate translates a key with optional attributes.
func Translate(c *gin.Context, key string, attributes ...map[string]string) string {
	var attr map[string]string
	if len(attributes) > 0 {
		attr = attributes[0] // Use the first attributes map if provided
	}

	locale, exists := c.Get(LocaleContextKey)
	if !exists || locale == "" {
		locale = GlobalLocaleConfig.DefaultLanguage // Fallback to default language
	}

	return TranslateKey(locale.(string), key, attr)
}

// TranslateKey translates a key into the specified locale
func TranslateKey(locale, key string, attributes map[string]string) string {
	translation := translations[locale][key]
	if translation == "" {
		translation = translations[GlobalLocaleConfig.DefaultLanguage][key] // Fallback to default language
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
	for _, l := range GlobalLocaleConfig.SupportedLanguages {
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
