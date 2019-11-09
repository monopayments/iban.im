package main

import (
	"github.com/monocash/iban.im/db"
	"github.com/monocash/iban.im/model"
)

func main() {
	d, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer d.Close()

	d.DropTableIfExists(&model.User{})
	d.DropTableIfExists(&model.Iban{})
	d.DropTableIfExists(&model.Group{})
	d.CreateTable(&model.User{})
	d.CreateTable(&model.Iban{})
	d.CreateTable(&model.Group{})
}
