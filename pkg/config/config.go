package config

import "os"

// AppConfig is our specefic application configs
type AppConfig struct {
	Name        string
	Version     string
	Port        string
	DatabaseURI string
}

//HttpConfig is our http configs
type HttpConfig struct {
	Content string
	Problem string
}

// Config is out global application's config
type Config struct {
	App  *AppConfig
	Http *HttpConfig
}

func NewConfig() *Config {
	return &Config{
		App: &AppConfig{
			Name:        env("APP_NAME", "Go app"),
			Version:     env("APP_VERSION", "0.0.1"),
			Port:        env("APP_PORT", "8000"),
			DatabaseURI: env("DATABASE_URI", ""),
		},
		Http: &HttpConfig{
			Content: env("HTTP_CONTENT", "application/json"),
			Problem: env("HTTP_PROBLEM", "application/problem+json"),
		},
	}
}

func env(key, defaultVal string) string {
	if value, exits := os.LookupEnv(key); exits {
		return value
	}
	return defaultVal
}
