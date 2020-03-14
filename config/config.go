package main

import (
	"github.com/jinzhu/configor"
)

type DBConfig struct {
	Name     string `env:"DBName" default:"ibanim"`
	Adapter  string `env:"DBAdapter" default:"postgres"`
	Host     string `env:"DBHost" default:"localhost"`
	Port     string `env:"DBPort" default:"5432"`
	User     string `env:"DBUser" default:"ibanim"`
	Password string `env:"DBPassword"`
}

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

var Config = struct {
	DB		DBConfig
	SMTP 	SMTPConfig
}{}

func init()  {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml"); err != nil {
		panic(err)
	}
}
