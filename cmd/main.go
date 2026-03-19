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
	"github.com/go-chi/chi"
)

func main() {
	//Load Env
	config := config.LoadConfig()

	//Server
	router := chi.NewRouter()
	server := http.Server{
		Handler: router,
		Addr:    ":" + config.Port,
	}
	go func() {
		log.Printf("Server Listenting on http:/localhost:%s\n", config.Port)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err.Error())
		}

	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server forced to shutdown: %s", err.Error())
	}
	log.Printf("Server Shutdowned\n")

}
