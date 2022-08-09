package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBhost     string `mapstructure:"DB_HOST"`
	DBport     string `mapstructure:"DB_PORT"`
	DBuser     string `mapstructure:"DB_USER"`
	DBpassword string `mapstructure:"DB_PASSWORD"`
	DBname     string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}

func InitDatabase() (*gorm.DB, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// dsn := "host=localhost user=postgres password=Masuk123 dbname=pjk_abt_mataram port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DBhost, config.DBuser, config.DBpassword, config.DBname, config.DBport)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	return db, nil

}
