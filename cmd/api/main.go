package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go/mini-s3/internal/config"
	"go/mini-s3/internal/http/router"
	"go/mini-s3/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env := config.Init()
	storage := storage.NewLocalStorage(env.StorageLocalPath)
	httpHandler := router.Init(storage, env)

	srv := &http.Server{
		Addr:              env.HTTPAddress,
		Handler:           httpHandler,
		ReadTimeout:       env.HTTPReadTimeout,
		WriteTimeout:      env.HTTPWriteTimeout,
		ReadHeaderTimeout: env.HTTPReadHeaderTimeout,
		IdleTimeout:       env.HTTPIdleTimeout,
	}

	go func() {
		log.Printf("server running on %v", env.HTTPAddress)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced shutdown: %v", err)
	}
}
