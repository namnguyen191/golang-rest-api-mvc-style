package config

import "os"

const (
	ENV_KEY_POSTGRES_USER     = "POSTGRES_USER"
	ENV_KEY_POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	ENV_KEY_POSTGRES_DB       = "POSTGRES_DB"
	ENV_KEY_JWT_SECRET        = "ENV_KEY_JWT_SECRET"
)

const (
	POSTGRES_USER     = "postgres"
	POSTGRES_PASSWORD = "postgres"
	POSTGRES_DB       = "themuzix"
	JWT_SECRET        = "98jnnvio9808e9kjfiasf9808"
)

// Initialize all necessary settings and configs
func SetInitialEnv() {
	os.Setenv(ENV_KEY_POSTGRES_USER, POSTGRES_USER)
	os.Setenv(ENV_KEY_POSTGRES_PASSWORD, POSTGRES_PASSWORD)
	os.Setenv(ENV_KEY_POSTGRES_DB, POSTGRES_DB)
	os.Setenv(ENV_KEY_JWT_SECRET, JWT_SECRET)
}
