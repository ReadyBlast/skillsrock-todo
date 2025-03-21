package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DbUser     string 
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	DbSslmode  string
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		viperConf := viper.New()
		viperConf.AddConfigPath("../internal/config")
		viperConf.SetConfigType("env")
		viperConf.SetConfigName(".env")

		if err := viperConf.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatalf("[ERROR]: Config file not found: %v", err)
			} else {
				log.Fatalf("[ERROR]: Error reading config file: %v", err)
			}
		} else {
			log.Println("[INFO]: Config file loaded successfully")
		}

		cfg = Config{
			DbUser:     viperConf.GetString("DB_USER"),
			DbPassword: viperConf.GetString("DB_PASSWORD"),
			DbHost:     viperConf.GetString("DB_HOST"),
			DbPort:     viperConf.GetString("DB_PORT"),
			DbName:     viperConf.GetString("DB_NAME"),
			DbSslmode:	viperConf.GetString("DB_SSLMODE"),
		}
	})
	return cfg
}
