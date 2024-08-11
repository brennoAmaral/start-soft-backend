package shared

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	envLoaded    bool = false
	envLoadMutex sync.Mutex
)

func loadEnv() {
	envLoadMutex.Lock()
	defer envLoadMutex.Unlock()

	if envLoaded {
		return
	}

	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	envLoaded = true
}

func GetEnv(key string) string {
	loadEnv()
	return os.Getenv(strings.ToUpper(key))
}
