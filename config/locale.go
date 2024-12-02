package config

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
	} else if contains(c.SupportedLanguages, c.DefaultLanguage) {
		c.SupportedLanguages = append(c.SupportedLanguages, c.DefaultLanguage)
	}

	// Set a default translation path if none is provided
	if len(c.TranslationPaths) == 0 {
		c.TranslationPaths = []string{"locales"}
	}
}

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
