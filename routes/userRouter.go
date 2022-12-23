package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmait/admin-finanzas/handlers"
	"github.com/osmait/admin-finanzas/server"
)

func UserRouters(s server.Server, r *mux.Router) {

	r.HandleFunc("/user", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handlers.GetUserById(s)).Methods(http.MethodGet)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
}
