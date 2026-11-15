package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DatabaseURL        string
	RedisAddr          string
	JWTSecret          string
	CORSOrigin         string
	SupabaseURL        string
	SupabaseServiceKey string
	SupabaseBucket     string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		Port:               getenv("PORT", "8080"),
		DatabaseURL:        databaseURL(),
		RedisAddr:          getenv("REDIS_ADDR", "localhost:6379"),
		JWTSecret:          getenv("JWT_SECRET", "dev-secret-ganti"),
		CORSOrigin:         getenv("CORS_ORIGIN", "https://group2.bccdev.id"),
		SupabaseURL:        getenv("SUPABASE_URL", ""),
		SupabaseServiceKey: getenvFirst([]string{"SUPABASE_SERVICE_ROLE_KEY", "SUPABASE_SERVICE_KEY"}, ""),
		SupabaseBucket:     getenvFirst([]string{"SUPABASE_STORAGE_BUCKET", "SUPABASE_BUCKET"}, "Lentera's Material"),
	}
}

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

func getenvFirst(keys []string, def string) string {
	for _, k := range keys {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return def
}