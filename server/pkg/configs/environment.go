package configs

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func GetEnv(key string) string {
    return os.Getenv(key)
}