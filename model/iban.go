package model

import (
	"time"

	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Iban : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type Iban struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Text      string `gorm:"type:varchar(100);not null"`
	Password  string
	Handle    string `gorm:"type:varchar(20);not null"`
	Active    bool
}
