package shared

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		return
	}
}

func GetEnv(key string) string {
	return os.Getenv(strings.ToUpper(key))
}
