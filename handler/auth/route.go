package auth

import (
	"net/http"

	"github.com/Blue-Onion/RestApi-Go/middleware"
	"github.com/go-chi/chi"
)

func AuthRoute(authHandler *Handler, middlewareHandler middleware.Handler) *chi.Mux {

	// User Routes
	userRoute := chi.NewRouter()
	userRoute.Post("/users", authHandler.HandleCreateUser)
	userRoute.Post("/login", authHandler.HandleLogin)
	userRoute.Post("/logOut", middlewareHandler.MiddlewareAuth(http.HandlerFunc(authHandler.HandleLogOut)))
	return userRoute
}
