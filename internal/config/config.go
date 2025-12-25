package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	LinkedInEmail    string
	LinkedInPassword string

	Headless         bool
	DailyConnectLimit int
}

func Load() *Config {
	cfg := &Config{
		LinkedInEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword: os.Getenv("LINKEDIN_PASSWORD"),
		Headless:         getEnvBool("HEADLESS", false),
		DailyConnectLimit: getEnvInt("DAILY_CONNECT_LIMIT", 20),
	}

	if cfg.LinkedInEmail == "" || cfg.LinkedInPassword == "" {
		log.Fatal("LINKEDIN_EMAIL or LINKEDIN_PASSWORD not set")
	}

	return cfg
}

func getEnvBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	parsed, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}
	return parsed
}

func getEnvInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return parsed
}
