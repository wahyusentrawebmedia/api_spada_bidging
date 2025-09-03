package database

import (
	"api/spada/internal/model"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.PostgresConfig{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
