package config

type FeatureConfig struct {
	EnableDB     bool
	EnableRedis  bool
	EnableLocale bool
}

func (f *FeatureConfig) BuildConfig() {
	if f.EnableLocale != true {
		f.EnableLocale = false
	}

	if f.EnableRedis != true {
		f.EnableRedis = false
	}

	if f.EnableLocale != true {
		f.EnableLocale = false
	}
}
