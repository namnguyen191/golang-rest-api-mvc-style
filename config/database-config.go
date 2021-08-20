package config

import (
	"fmt"
	"os"

	"github.com/namnguyen191/themuzix-golang-rest-api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Create a connetion to the database
func SetupDatabaseConnection() *gorm.DB {
	var (
		dbName     = os.Getenv(ENV_KEY_POSTGRES_DB)
		dbUser     = os.Getenv(ENV_KEY_POSTGRES_USER)
		dbPass     = os.Getenv(ENV_KEY_POSTGRES_PASSWORD)
		dbHost     = "localhost"
		dbPort     = 5432
		dbSslMode  = "disable"
		dbTimeZone = "UTC"
	)

	if dbName == "" || dbUser == "" || dbPass == "" {
		panic("missing env variable")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%d sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbPort, dbSslMode, dbTimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect to db")
	}

	// Migrate DB model
	err = db.AutoMigrate(&entity.Artist{}, &entity.Song{}, &entity.User{})
	if err != nil {
		panic("fail to auto migrate tables")
	}

	return db
}

// close database connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Printf("fail to close db connetion: %v", err)
		return
	}

	dbSQL.Close()
}
