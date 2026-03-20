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
	//
	router.Get("/health",handler.Health)
	stop:=make(chan os.Signal,1)
	go func() {
		log.Print("Listening on http:/localhost:3480")
		err:=server.ListenAndServe()
		if err!=nil{
			log.Fatal("error occured")
		}
	}()
	<-stop
	signal.Notify(stop,os.Interrupt,syscall.SIGTERM)
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()
	err:=server.Shutdown(ctx)
	if err!=nil{
		log.Fatal("Error occured in Shutdown")
	}
	log.Fatal("Server Shutdown gracefully")

}
