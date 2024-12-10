package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
		Encoding string
	}
}

func LoadConfig() *Config {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Errorf("error loading config: %v", err)
		return nil
	}

	config := &Config{}

	config.DB.Host = os.Getenv("DB_HOST")
	config.DB.Port = os.Getenv("DB_PORT")
	config.DB.User = os.Getenv("DB_USER")
	config.DB.Password = os.Getenv("DB_PASSWORD")
	config.DB.DBName = os.Getenv("DB_NAME")
	config.DB.SSLMode = os.Getenv("DB_SSL_MODE")
	config.DB.Encoding = os.Getenv("DB_ENCODING")

	return config
}

func (config *Config) GetDBConnString() string {
	return "postgres" + "://" +
		config.DB.User + ":" +
		config.DB.Password + "@" +
		config.DB.Host + ":" +
		config.DB.Port + "/"
}

func (config *Config) GetMigrationDBConnString() string {
	return config.GetDBConnString() + "?" +
		config.DB.SSLMode + "&" +
		config.DB.Encoding
}

func (config *Config) GetMusicDBConnString() string {
	return config.GetDBConnString() +
		config.DB.DBName + "?" +
		config.DB.SSLMode + "&" +
		config.DB.Encoding
}
