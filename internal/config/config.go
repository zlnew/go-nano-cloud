// Package config
package config

import (
	"os"
	"strconv"
	"time"
)

type BaseEnv struct {
	HTTPAddress           string
	HTTPReadTimeout       time.Duration
	HTTPWriteTimeout      time.Duration
	HTTPReadHeaderTimeout time.Duration
	HTTPIdleTimeout       time.Duration
	StorageLocalPath      string
}

func Init() *BaseEnv {
	return &BaseEnv{
		HTTPAddress:           getEnv("HTTP_ADDRESS", ":3000"),
		HTTPReadTimeout:       getDurationEnv("HTTP_READ_TIMEOUT", 5),
		HTTPWriteTimeout:      getDurationEnv("HTTP_WRITE_TIMEOUT", 10),
		HTTPReadHeaderTimeout: getDurationEnv("HTTP_READ_HEADER_TIMEOUT", 2),
		HTTPIdleTimeout:       getDurationEnv("HTTP_IDLE_TIMEOUT", 60),
		StorageLocalPath:      getEnv("STORAGE_LOCAL_PATH", "uploads"),
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
