package config

import (
	"log"
	"os"
)

type Config struct {
	App struct {
		Port string
	}
	DB struct {
		URL string
	}
	MQ struct {
		URL string
	}
}

func Load() *Config {
	cfg := &Config{}
	cfg.App.Port = getEnv("KEEAUDIT_PORT", "8080")
	cfg.DB.URL = mustEnv("KEEAUDIT_DBURL")
	cfg.MQ.URL = mustEnv("KEEAUDIT_MQURL")
	return cfg
}

func getEnv(key string, fallback string) string {
	s := os.Getenv(key)
	if s == "" {
		return fallback
	}
	return s
}

func mustEnv(key string) string {
	s := os.Getenv(key)
	if s == "" {
		log.Fatalf("missing required env var: %s", key)
	}
	return s
}
