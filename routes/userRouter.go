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
	r.HandleFunc("/user/confirmed/{token}", handlers.UserTokenAuth(s)).Methods(http.MethodGet)

	// Perfil
	r.HandleFunc("/perfil", handlers.Perfil(s)).Methods(http.MethodGet)

	// Account
	r.HandleFunc("/account", handlers.InsertAccount(s)).Methods(http.MethodPost)
	r.HandleFunc("/account", handlers.GetAccounts(s)).Methods(http.MethodGet)
	r.HandleFunc("/account/{id}", handlers.DeleteAccount(s)).Methods(http.MethodDelete)

	// Transactions
	r.HandleFunc("/transation", handlers.InsertTransaction(s)).Methods(http.MethodPost)
	r.HandleFunc("/transation/general", handlers.GetAllTransaction(s)).Methods(http.MethodGet)
	r.HandleFunc("/transation/{id}", handlers.GetTransaction(s)).Methods(http.MethodGet)
	r.HandleFunc("/transation/{id}", handlers.UpdateTransaction(s)).Methods(http.MethodPut)

	r.HandleFunc("/transation/{id}", handlers.DeleteTransaction(s)).Methods(http.MethodDelete)
	// Balances
	r.HandleFunc("/balance", handlers.GetBalace(s)).Methods(http.MethodGet)

}
