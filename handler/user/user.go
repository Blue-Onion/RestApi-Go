package user

import (
	"encoding/json"
	"net/http"
	"github.com/Blue-Onion/RestApi-Go/handler"
	"github.com/Blue-Onion/RestApi-Go/model"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	param := model.CreateUser{}
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&param)
	if err!=nil{
		handler.RespondWithError(w,400,"Error in Parsing Json")
		return
	}

	handler.RespondWithJson(w, 201, param)
}
