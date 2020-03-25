package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/monocash/iban.im/config"
	"strings"
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

// Check Handle before create or update = must be add as index to db
func (iban *Iban) CheckHandle() (exist bool) {
	var ibans []Iban
	config.DB.Where("owner_id = ? & handle = ?",iban.OwnerID,iban.Handle).Find(&ibans)
	for _, tmp := range ibans {
		if iban.Handle == tmp.Handle && iban.IbanID != tmp.IbanID {
			exist = true
		}
	}
	return
}

// BeforeSave Callback
func (iban  *Iban) BeforeSave(tx *gorm.DB) (err error) {
	if iban.CheckHandle() {
		err = fmt.Errorf("handle already exist")
	}
	return
}

func (iban *Iban) Validate(db *gorm.DB)  {
	if strings.TrimSpace(iban.Text) == "" {
		db.AddError(fmt.Errorf("you have to provide IBAN"))
	}else if strings.TrimSpace(iban.Handle) == "" {
		db.AddError(fmt.Errorf("you have to provide handle"))
	}else if iban.IsPrivate && strings.TrimSpace(iban.Password) == "" {
		db.AddError(fmt.Errorf("you have to provide password"))
	}
}

func GetIbanById(id uint) Iban  {
	iban := Iban{}
	config.DB.First(&iban,id)
	return iban
}