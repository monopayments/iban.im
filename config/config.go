package config

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

type AppConfig struct {
	Port  uint `default:"7000" env:"PORT"`
	Env   string `default:"localhost" env:"ENV"`
}

var Config = struct {
	DB   DBConfig
	SMTP SMTPConfig
	APP  AppConfig
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml","config/application.yml"); err != nil {
		panic(err)
	}
}
