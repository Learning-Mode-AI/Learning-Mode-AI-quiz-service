package config

import (
	"fmt"
	"os"
)

var (
	RedisHost string
	AIHost string
)

func InitConfig() {
	env := os.Getenv("ENVIRONMENT")
	if env == "local" {
		RedisHost = "localhost:6379"
		AIHost = "localhost:8082"
		fmt.Println("Running in local mode")
	} else {
		RedisHost = "redis:6379"
		AIHost = "http://ai-service:8082"
		fmt.Println("Running in Docker mode")
	}
}
