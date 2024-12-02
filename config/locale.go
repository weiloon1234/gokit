package config

import "github.com/weiloon1234/gokit/utils"

type LocaleConfig struct {
	DefaultLanguage    string
	SupportedLanguages []string
	TranslationPaths   []string // Paths to search for project-level locale files
}

func (c *LocaleConfig) BuildConfig() {
	// Set default language if not provided
	if c.DefaultLanguage == "" {
		c.DefaultLanguage = "en"
	}

	// Ensure SupportedLanguages includes the default language
	if len(c.SupportedLanguages) == 0 {
		c.SupportedLanguages = []string{"en"} // Default to English
	} else if !utils.Contains(c.SupportedLanguages, c.DefaultLanguage) {
		c.SupportedLanguages = append(c.SupportedLanguages, c.DefaultLanguage)
	}

	// Set a default translation path if none is provided
	if len(c.TranslationPaths) == 0 {
		c.TranslationPaths = []string{"locales"}
	}
}
