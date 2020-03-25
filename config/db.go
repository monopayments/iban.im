package config

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/monocash/iban.im/model"
	"github.com/qor/validations"

	// _ "github.com/jinzhu/gorm/dialects/sqlite" TODO - disabled for compile time issue
	"os"
	"time"
)

// global DB variable => TODO repository pattern
var DB *gorm.DB

func init() {
	var err error
	if Config.App.Env == "gitpod" && Config.Db.Adapter == "postgres" {
		if err = os.Unsetenv("PGHOSTADDR"); err != nil {
			panic(err)
		}
	}

	var connStr string
	adapter := Config.Db.Adapter
	if adapter == "mysql" {
		connStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", Config.Db.User, Config.Db.Password, Config.Db.Host, Config.Db.Port, Config.Db.Name)
	} else if adapter == "postgres" {
		connStr = fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", Config.Db.User, Config.Db.Password, Config.Db.Host, Config.Db.Name)
	} else if adapter == "sqlite3" || adapter == "sqlite" {
		connStr = fmt.Sprintf("%v/%v", os.TempDir(), Config.Db.Name)
	} else {
		panic(errors.New("your database is not supported"))
	}

	DB, err = gorm.Open(adapter, connStr)
	if err != nil {
		panic(err)
	}
	validations.RegisterCallbacks(DB)
	DB.LogMode(Config.App.Debug)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(30)
	DB.DB().SetConnMaxLifetime(time.Second * 60)

	DB.AutoMigrate(&model.User{},&model.Iban{},&model.Group{})

	// TODO ping control for mysql
	if adapter == "mysql" {
		go checkPing()
	}
}

func checkPing() {
	for {
		time.Sleep(time.Second * 15)
		DB.DB().Ping()
	}
}
