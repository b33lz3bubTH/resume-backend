package config

import (
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	Port           string
	DatabaseURL    string
	RootKey        string
	Environment    string
	SuperadminEmail string
}

func Load() *Config {
	databaseURL := getEnv("DATABASE_URL", "")
	if databaseURL == "" {
		dbPass := getEnv("DB_PASS", "")
		encodedPass := url.QueryEscape(dbPass)
		databaseURL = fmt.Sprintf("postgresql://postgres.rrytjodvjedesjifxwpp:%s@aws-1-ap-south-1.pooler.supabase.com:6543/postgres?sslmode=require&pgbouncer=true", encodedPass)
	}
	
	return &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    databaseURL,
		RootKey:        getEnv("ROOT_KEY", "abcd123"),
		Environment:    getEnv("ENV", "development"),
		SuperadminEmail: getEnv("SUPERADMIN_EMAIL", "souravsunju@gmail.com"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

