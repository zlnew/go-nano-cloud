package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go/mini-s3/internal/config"
	router "go/mini-s3/internal/http"
	"go/mini-s3/internal/storage"
)

func main() {
	godotenv.Load()
	conf := config.Init()
	s := storage.NewLocalStorage(conf.StorageLocalPath)
	r := router.Init(s)

	srv := &http.Server{
		Addr:              conf.HTTPAddress,
		Handler:           r,
		ReadTimeout:       conf.HTTPReadTimeout,
		WriteTimeout:      conf.HTTPWriteTimeout,
		ReadHeaderTimeout: conf.HTTPReadHeaderTimeout,
		IdleTimeout:       conf.HTTPIdleTimeout,
	}

	go func() {
		log.Printf("Server running on %v", conf.HTTPAddress)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}
}
