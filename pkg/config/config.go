package config

import (
	"log"
	"net/url"
	"os"
)

type Config struct {
	Port           string
	DatabaseURL    string
	RootKey        string
	Environment    string
	SuperadminEmail string
	OpenRouterKey  string
	OpenRouterModel string
}

func Load() *Config {
	databaseURL := getEnv("DATABASE_URL", "")
	if databaseURL == "" {
		dbPass := getEnv("DB_PASS", "")
		encodedPass := url.QueryEscape(dbPass)
		databaseURL = "postgresql://postgres.rrytjodvjedesjifxwpp:" + encodedPass + "@aws-1-ap-south-1.pooler.supabase.com:6543/postgres"
		
		log.Printf("DEBUG: DB_PASS from env: %s", dbPass)
		log.Printf("DEBUG: Encoded password: %s", encodedPass)
		log.Printf("DEBUG: Full database URL: %s", databaseURL)
	} else {
		log.Printf("DEBUG: Using DATABASE_URL from environment (password hidden)")
		log.Printf("DEBUG: Database URL (first 50 chars): %.50s...", databaseURL)
	}
	
	return &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    databaseURL,
		RootKey:        getEnv("ROOT_KEY", "abcd123"),
		Environment:    getEnv("ENV", "development"),
		SuperadminEmail: getEnv("SUPERADMIN_EMAIL", "souravsunju@gmail.com"),
		OpenRouterKey:  getEnv("OPENROUTER_API_KEY", "sk-or-v1-"),
		OpenRouterModel: getEnv("OPENROUTER_MODEL", "allenai/olmo-3.1-32b-think:free"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

