package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Blue-Onion/RestApi-Go/config"
	"github.com/Blue-Onion/RestApi-Go/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	//Load Env
	config := config.LoadConfig()

	//Server
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	server := http.Server{
		Handler: router,
		Addr:    ":" + config.Port,
	}
	//
	router.Get("/health",handler.Health)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Listening on http://localhost:%s", config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error occurred: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error occurred in Shutdown: %v", err)
	}
	log.Println("Server Shutdown gracefully")
}
