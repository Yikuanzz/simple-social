package conf

import (
	"os"
	"strconv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	valueStr := getEnv(key, "")
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return fallback
	}
	return value
}
