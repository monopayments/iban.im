package db

import (
	"github.com/jinzhu/gorm"
	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB *grom.DB
type DB struct {
	*gorm.DB
}

// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=ibanim dbname=ibanim password=ibanim sslmode=disable")
	// db, err := gorm.Open("postgres", "host=host.docker.internal port=5432 user=ibanim dbname=ibanim password=ibanim sslmode=disable")

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}
