package config

type FeatureConfig struct {
	EnableDB     *bool
	EnableRedis  *bool
	EnableLocale *bool
}

func (f *FeatureConfig) BuildConfig() {
	if f.EnableLocale == nil {
		defaultValue := true // Set your default value here
		f.EnableLocale = &defaultValue
	}

	if f.EnableDB == nil {
		defaultValue := false
		f.EnableDB = &defaultValue
	}

	if f.EnableRedis == nil {
		defaultValue := false
		f.EnableRedis = &defaultValue
	}
}
