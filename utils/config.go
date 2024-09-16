package utils

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	Database string `envconfig:"DB_DATABASE"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
}

type AppConfig struct {
	DBConfig DBConfig
}

func LoadConfig() {
	_, fileStr, _, _ := runtime.Caller(0)

	rootPath := filepath.Join(filepath.Dir(fileStr), "../")

	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		log.Fatal("error loading .env ", err)
	}

	loadConfig()
}

var config *AppConfig = &AppConfig{}

func loadConfig() {

	err := envconfig.Process("", config)
	if err != nil {
		log.Fatal("Error loadconfig", err.Error())
	}
}

func Config() *AppConfig {
	if config == nil {
		LoadConfig()
	}

	return config
}
