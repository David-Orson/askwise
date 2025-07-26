package utils

import (
	"log"
	"os"

	"askwise.com/m/v2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Missing DATABASE_URL environment variable")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	var version string
	if err := db.Raw("SELECT version()").Scan(&version).Error; err != nil {
		log.Fatalf("Failed to query version: %v", err)
	}

	log.Println("✅ Connected to Postgres version:", version)

	db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Document{},
		&models.Message{},
	)
	return db
}
