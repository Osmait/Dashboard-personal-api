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

	// Income
	r.HandleFunc("/transation", handlers.InsertIncome(s)).Methods(http.MethodPost)
	r.HandleFunc("/transation/general", handlers.GetAllTransaction(s)).Methods(http.MethodGet)
	r.HandleFunc("/transation/{id}", handlers.GetIncome(s)).Methods(http.MethodGet)
	r.HandleFunc("/transation/{id}", handlers.DeleteIncome(s)).Methods(http.MethodDelete)

	// // Bills
	// r.HandleFunc("/bill", handlers.InsertBill(s)).Methods(http.MethodPost)
	// r.HandleFunc("/bill/{id}", handlers.GetBills(s)).Methods(http.MethodGet)
	// r.HandleFunc("/bill/{id}", handlers.DeleteBill(s)).Methods(http.MethodDelete)

	// // IncomeAndBill
	// r.HandleFunc("/icomeandbill/{id}", handlers.IncomeAndBill(s)).Methods(http.MethodGet)

}
