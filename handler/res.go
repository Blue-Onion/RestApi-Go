package handler

import (
	"encoding/json"
	"log"
	"net/http"


)



func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to parse Json")
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-type", "Application/Json")
	w.WriteHeader(code)
	w.Write(data)

}
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error", msg)
	}
	type response struct {
		Msg string `json:"Error"`
	}
	RespondWithJson(w, code, response{Msg: msg})
}
func Health(w http.ResponseWriter, r *http.Request) {
	type res struct {
		Health string
	}
	response := res{
		Health: "ok",
	}
	RespondWithJson(w, 200, response)
}
