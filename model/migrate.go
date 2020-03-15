package model

import "github.com/monocash/iban.im/config"

func init()  {
	config.DB.AutoMigrate(&User{},&Iban{},&Group{})
}
