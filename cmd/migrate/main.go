package main

import (
	"event-bali/internal/config"
	"event-bali/internal/database"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.PostgresDSN(cfg.DBDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")
}
