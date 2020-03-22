package model

import (
	"time"

	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

// Iban : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type Iban struct {
	IbanID    uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Text      string `gorm:"type:varchar(100);not null"`
	Password  string
	Handle    string `gorm:"type:varchar(20);not null"`
	Active    bool
	IsPrivate bool
	OwnerID   uint
	OwnerType string
}

// HashPassword : hashing the password
func (iban *Iban) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(iban.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	iban.Password = string(hash)
}

// ComparePassword : compare the password
func (iban *Iban) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(iban.Password), []byte(password))

	if err != nil {
		return false
	}

	return true
}
