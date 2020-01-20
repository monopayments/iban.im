package db

import (
	"os"

	"github.com/jinzhu/gorm"
	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB *grom.DB
type DB struct {
	*gorm.DB
}

var connStrMap = map[string]string {
	"localhost" : "host=localhost port=5432 user=ibanim dbname=ibanim password=ibanim sslmode=disable",
	"docker" : "host=host.docker.internal port=5432 user=ibanim dbname=ibanim password=ibanim sslmode=disable",
	"gitpod" : "host=localhost port=5432 user=ibanim dbname=ibanim password=ibanim sslmode=disable",
}

// ConnectDB : connecting DB
func ConnectDB(env string) (*DB, error) {
	if env == "gitpod" {
		os.Unsetenv("PGHOSTADDR")
	}
	db, err := gorm.Open("postgres", connStrMap[env])
	// db, err := gorm.Open("postgres", "host=host.docker.internal port=5432 user=ibanim dbname=ibanim password=ibanim sslmode=disable")

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}
