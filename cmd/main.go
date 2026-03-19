package main

import (
	"log"
	"net/http"

	"github.com/Blue-Onion/RestApi-Go/config"
	"github.com/go-chi/chi"
)

func main() {
	//Load Env
	config:=config.LoadConfig()

	//Server
	router:=chi.NewRouter()
	server:=http.Server{
		Handler: router,
		Addr: ":"+config.Port,
	}
	go func() {
		err:=server.ListenAndServe()
		if err!=nil{
			log.Fatal(err.Error())
		}

	}()
}
