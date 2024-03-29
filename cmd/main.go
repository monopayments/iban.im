package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/monopayments/iban.im/model"
	"log"
)

var DBOld *gorm.DB
var DBNew *gorm.DB

func init() {
	var err error
	DBOld, err = gorm.Open("mysql", "iban_p8nPjkfKO0M:T1os3vVUiQlZeEw@/iban_pushecommerce_com?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	DBNew, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ibanim password=Ahmety61+- sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	DBNew.AutoMigrate(&model.User{}, &model.Group{}, &model.Iban{})
}

func main() {
	migrateUser()
	migrateIban()
	migrateGroups()
}

func migrateUser() {
	var users []model.User
	DBOld.Limit(-1).Find(&users)
	for _, user := range users {
		DBNew.Create(&user)
	}
}

func migrateIban() {
	var ibans []model.Iban
	DBOld.Limit(-1).Find(&ibans)
	for _, iban := range ibans {
		DBNew.Create(&iban)
	}
}

func migrateGroups() {
	var groups []model.Group
	DBOld.Limit(-1).Find(&groups)
	for _, group := range groups {
		DBNew.Create(&group)
	}
}
