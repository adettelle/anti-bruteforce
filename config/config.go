package config

import "os"

const (
	defaultPort = "8080"
)

type Config struct {
	Port string
}

func New() *Config {
	cfg := Config{
		Port: getEnvOrDefault("PORT", defaultPort),
	}
	return &cfg
}

func getEnvOrDefault(envName string, defaultVal string) string {
	res := os.Getenv(envName)
	if res == "" {
		return defaultVal
	}
	return res
}
