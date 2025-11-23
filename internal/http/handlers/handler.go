// Package handlers
package handlers

import "go/mini-s3/internal/storage"

type StorageHandler struct {
	Storage storage.Storage
}

func NewStorageHandler(s storage.Storage) *StorageHandler {
	return &StorageHandler{Storage: s}
}
