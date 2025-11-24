// Package handlers
package handlers

import (
	"go/nano-cloud/internal/config"
	"go/nano-cloud/internal/storage"
)

type StorageHandler struct {
	Storage storage.Storage
	Env     *config.BaseEnv
}

func NewStorageHandler(s storage.Storage, env *config.BaseEnv) *StorageHandler {
	return &StorageHandler{Storage: s, Env: env}
}
