package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/jinzhu/gorm"
	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Model : gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Iban : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type Iban struct {
	gorm.Model
	Text      string `gorm:"type:varchar(100);not null"`
	Password  string
	Handle	  string `gorm:"type:varchar(20);not null"`
}

