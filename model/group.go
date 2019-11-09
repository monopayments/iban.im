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

// Group : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type Group struct {
	gorm.Model
	GroupName      string `gorm:"type:varchar(100);not null"`
	Handle	  string `gorm:"type:varchar(20);not null"`
}

