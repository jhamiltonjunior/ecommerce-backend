package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// America/Sao_Paulo
func OpenDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATA_SOURCE_NAME")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
