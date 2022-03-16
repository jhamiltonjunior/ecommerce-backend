package configs

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// America/Sao_Paulo
func OpenDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DATA_SOURCE_NAME"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}


// , &gorm.Config{})