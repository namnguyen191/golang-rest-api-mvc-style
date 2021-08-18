package config

import "os"

const (
	ENV_KEY_POSTGRES_USER     = "POSTGRES_USER"
	ENV_KEY_POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	ENV_KEY_POSTGRES_DB       = "POSTGRES_DB"
)

const (
	POSTGRES_USER     = "postgres"
	POSTGRES_PASSWORD = "postgres"
	POSTGRES_DB       = "themuzix"
)

// Intialize all neccessary settings and configs
func SetIntialEnv() {
	os.Setenv(ENV_KEY_POSTGRES_USER, POSTGRES_USER)
	os.Setenv(ENV_KEY_POSTGRES_PASSWORD, POSTGRES_PASSWORD)
	os.Setenv(ENV_KEY_POSTGRES_DB, POSTGRES_DB)
}
