package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME"`
	ServicePort string `envconfig:"SERVICE_PORT"`
	ServiceEnv  string `envconfig:"SERVICE_ENV"`

	DbMaster   string `envconfig:"RDS_MASTER"`
	DbReplica  string `envconfig:"RDS_REPLICA"`
	DbUsername string `envconfig:"RDS_USERNAME"`
	DbPassword string `envconfig:"RDS_PASSWORD"`
	DbName     string `envconfig:"RDS_DBNAME"`
	DbPort     string `envconfig:"RDS_PORT"`

	JwtSecretKey string `envconfig:"JWT_SECRET_KEY"`
}

func LoadEnv(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}
