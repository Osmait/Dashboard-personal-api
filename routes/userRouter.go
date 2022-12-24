package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmait/admin-finanzas/handlers"
	"github.com/osmait/admin-finanzas/middleware"
	"github.com/osmait/admin-finanzas/server"
)

func UserRouters(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	// User
	r.HandleFunc("/user", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handlers.GetUserById(s)).Methods(http.MethodGet)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	// Account
	r.HandleFunc("/account", handlers.InsertAccount(s)).Methods(http.MethodPost)
	r.HandleFunc("/account", handlers.GetAccounts(s)).Methods(http.MethodGet)
	r.HandleFunc("/account/{id}", handlers.DeleteAccount(s)).Methods(http.MethodDelete)

	// Bills
	r.HandleFunc("/income", handlers.InsertIncome(s)).Methods(http.MethodPost)
	r.HandleFunc("/income/{id}", handlers.GetIncome(s)).Methods(http.MethodGet)
	r.HandleFunc("/income/{id}", handlers.DeleteIncome(s)).Methods(http.MethodDelete)

}
