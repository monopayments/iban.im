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
	Password string `env:"DBPassword" default:"ibanim"`
}

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type AppConfig struct {
	Port  uint   `default:"7000" env:"PORT"`
	Env   string `default:"localhost" env:"ENV"`
	Debug bool   `default:"false" env:"DEBUG"`

	// Session Timeout - minutes
	Timeout uint `default:"60"`

	// Auth Max Refresh - minutes
	MaxRefresh uint `default:"60"`
	
	// Auth Key
	Key string `default:"12345678" env:"AUTH_KEY"`

	// Realm name to display to the user.
	Realm string `default:"ibanim zone"`
}

var Config = struct {
	Db   DBConfig
	Smtp SMTPConfig
	App  AppConfig
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml", "config/application.yml"); err != nil {
		panic(err)
	}
}
