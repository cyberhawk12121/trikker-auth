package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/trikker/auth/internal/api"
	"github.com/trikker/auth/internal/db"
	"github.com/trikker/auth/internal/repository"
	"github.com/trikker/auth/internal/service"
)

func main() {
	cfg, err := db.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load configurations: %v", err)
	}

	database, err := db.Connect(&cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.CloseDB()

	userRepo := repository.NewSQLUserRepository(database)
	authService := service.NewAuthService(userRepo, db.JWTSecret())

	router := api.SetupRouter(authService)

	srv := &http.Server{
		Addr:    db.ServerAddr(),
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-quit
	log.Println("Received interrupt signal, shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}