package router

import (
	"net/http"

	"go/mini-s3/internal/config"
	"go/mini-s3/internal/http/handlers"
	"go/mini-s3/internal/http/middleware"
	"go/mini-s3/internal/storage"
)

func Init(s storage.Storage, env *config.BaseEnv) http.Handler {
	r := New()

	withMiddleware := middleware.Chain(
		middleware.Logger,
		middleware.Recoverer,
	)

	storageHandler := handlers.NewStorageHandler(s, env)

	r.Post("/upload", withMiddleware(http.HandlerFunc(storageHandler.Upload)))
	r.Get("/files", withMiddleware(http.HandlerFunc(storageHandler.List)))
	r.Get("/files/", withMiddleware(http.HandlerFunc(storageHandler.Download)))
	r.Delete("/delete/", withMiddleware(http.HandlerFunc(storageHandler.Delete)))

	return r.Handler()
}
