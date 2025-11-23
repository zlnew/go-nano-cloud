// Package router
package router

import (
	"net/http"

	"go/mini-s3/internal/config"
	"go/mini-s3/internal/http/handlers"
	"go/mini-s3/internal/http/middleware"
	"go/mini-s3/internal/storage"
)

func Init(s storage.Storage, env *config.BaseEnv) http.Handler {
	mux := http.NewServeMux()

	withMiddleware := middleware.Chain(
		middleware.Logger,
		middleware.Recoverer,
	)

	storageHandler := handlers.NewStorageHandler(s, env)

	mux.Handle("/ping", withMiddleware(http.HandlerFunc(handlers.Ping)))
	mux.Handle("/upload", withMiddleware(http.HandlerFunc(storageHandler.Upload)))
	mux.Handle("/files", withMiddleware(http.HandlerFunc(storageHandler.List)))
	mux.Handle("/files/", withMiddleware(http.HandlerFunc(storageHandler.Download)))
	mux.Handle("/delete/", withMiddleware(http.HandlerFunc(storageHandler.Delete)))

	return mux
}
