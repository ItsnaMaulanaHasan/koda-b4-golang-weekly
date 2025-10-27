package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadDefaultEnv(key string, defaultValue string) string {
	godotenv.Load()
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return value
}
