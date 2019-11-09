package model

import (
	"time"

	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Group : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type Group struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	GroupName string `gorm:"type:varchar(100);not null"`
	Handle    string `gorm:"type:varchar(20);not null"`
}
