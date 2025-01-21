package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var projectEnv string

func PanicIfNotDev() {
	if !IsDev() {
		panic("this service is only for development, as all the methods are mocked")
	}
}

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	projectEnv = os.Getenv("ENV")
}

func ProjectEnv() string {
	return projectEnv
}

func IsDev() bool {
	return ProjectEnv() == "dev"
}
