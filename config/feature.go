package config

type Features struct {
	EnableDB     *bool
	EnableRedis  *bool
	EnableLocale *bool
}

func (f *Features) BuildConfig() {
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
