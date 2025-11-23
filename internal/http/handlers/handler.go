// Package handlers
package handlers

import (
	"go/mini-s3/internal/config"
	"go/mini-s3/internal/storage"
)

type StorageHandler struct {
	Storage storage.Storage
	Env     *config.BaseEnv
}

func NewStorageHandler(s storage.Storage, env *config.BaseEnv) *StorageHandler {
	return &StorageHandler{Storage: s, Env: env}
}
