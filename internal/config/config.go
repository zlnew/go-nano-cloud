// Package config
package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type BaseEnv struct {
	HTTPAddress           string
	HTTPReadTimeout       time.Duration
	HTTPWriteTimeout      time.Duration
	HTTPReadHeaderTimeout time.Duration
	HTTPIdleTimeout       time.Duration
	StorageLocalPath      string
	MaxRequestBodySize    int64
	MaxMultipartMemory    int64
	APIKey                string
}

func Init() *BaseEnv {
	godotenv.Load()

	return &BaseEnv{
		HTTPAddress:           getEnv("HTTP_ADDRESS", ":3000"),
		HTTPReadTimeout:       getDurationEnv("HTTP_READ_TIMEOUT", 5),
		HTTPWriteTimeout:      getDurationEnv("HTTP_WRITE_TIMEOUT", 10),
		HTTPReadHeaderTimeout: getDurationEnv("HTTP_READ_HEADER_TIMEOUT", 2),
		HTTPIdleTimeout:       getDurationEnv("HTTP_IDLE_TIMEOUT", 60),
		StorageLocalPath:      getEnv("STORAGE_LOCAL_PATH", "uploads"),
		MaxRequestBodySize:    getInt64Env("MAX_REQUEST_BODY_SIZE", 20) * 1024 * 1024,
		MaxMultipartMemory:    getInt64Env("MAX_MULTIPART_MEMORY", 8) * 1024 * 1024,
		APIKey:                getRequiredEnv("API_KEY"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getDurationEnv(key string, fallbackSeconds int) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return time.Duration(fallbackSeconds) * time.Second
	}

	sec, err := strconv.Atoi(val)
	if err != nil {
		return time.Duration(fallbackSeconds) * time.Second
	}

	return time.Duration(sec) * time.Second
}

func getInt64Env(key string, fallback int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return fallback
	}

	return num
}

func getRequiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s is required", key)
	}

	return val
}
