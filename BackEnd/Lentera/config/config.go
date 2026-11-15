package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	RedisAddr   string
	JWTSecret   string
	CORSOrigin  string
}

func Load() Config {
	// muat .env ke environment proses (buat go run lokal).
	// di Docker file .env gak ada di image → error diabaikan, env dari compose yang dipakai.
	_ = godotenv.Load()

	return Config{
		Port:        getenv("PORT", "8080"),
		DatabaseURL: databaseURL(),
		RedisAddr:   getenv("REDIS_ADDR", "localhost:6379"),
		JWTSecret:   getenv("JWT_SECRET", "dev-secret-ganti"),
		CORSOrigin:  getenv("CORS_ORIGIN", "http://localhost:3000"),
	}
}

// pakai DATABASE_URL kalau di-set, kalau tidak rakit dari POSTGRES_* + DB_*
func databaseURL() string {
	if v := os.Getenv("DATABASE_URL"); v != "" {
		return v
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		getenv("DB_HOST", "localhost"),
		getenv("DB_PORT", "5432"),
		getenv("POSTGRES_USER", "lentera"),
		getenv("POSTGRES_PASSWORD", "lentera123"),
		getenv("POSTGRES_DB", "lentera"),
		getenv("DB_SSLMODE", "disable"),
	)
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
