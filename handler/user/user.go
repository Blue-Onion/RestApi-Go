package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Blue-Onion/RestApi-Go/handler"
	"github.com/Blue-Onion/RestApi-Go/internal/database"
	"github.com/Blue-Onion/RestApi-Go/model"
	"github.com/Blue-Onion/RestApi-Go/utlis"
	"github.com/google/uuid"
)

type Handler struct {
	Repo database.UserRepository
}

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	params := model.AutheticateUser{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		handler.RespondWithError(w, 400, "Error in Parsing Json")
		return
	}
	user,err:=h.Repo.GetUserByEmail(r.Context(),params.Email)
	if err!=nil{
		handler.RespondWithError(w, 400, err.Error())
		return
		
	}
	isValid:=utlis.CheckPassword(user.Password,params.Password)
	if !isValid{
		handler.RespondWithError(w, 400, "Incorrect Password")
		return

	}
	handler.RespondWithJson(w,200,user)
}

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	param := model.CreateUser{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&param)
	if err != nil {
		handler.RespondWithError(w, 400, "Error in Parsing Json")
		return
	}
	hashPass, err := utlis.HashPassword(param.Password)
	if err != nil {
		handler.RespondWithError(w, 400, "Error in Hashing Password")
		return
	}
	user, err := h.Repo.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      param.Name,
		Email:     param.Email,
		Password:  hashPass,
		Createdat: time.Now(),
		Updatedat: time.Now(),
	})
	if err != nil {
		handler.RespondWithError(w, 500, "Couldn't create user")
		return
	}

	handler.RespondWithJson(w, 201, user)
}
