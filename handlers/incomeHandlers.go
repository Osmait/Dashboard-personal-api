package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/osmait/admin-finanzas/helpers"
	"github.com/osmait/admin-finanzas/models"
	"github.com/osmait/admin-finanzas/repository"
	"github.com/osmait/admin-finanzas/server"
	"github.com/segmentio/ksuid"
)

func InsertIncome(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := helpers.DecodeJwt(w, r, s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*models.AppClaims)

		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return

		}
		var income = models.Transaction{}
		err = json.NewDecoder(r.Body).Decode(&income)
		fmt.Println(income)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		income.Id = id.String()
		income.UserId = claims.UserId
		err = repository.InsertIncome(r.Context(), &income)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(income)

	}
}
func GetIncome(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		date1 := r.URL.Query().Get("date1")
		date2 := r.URL.Query().Get("date2")

		if date1 == "" || date2 == "" {
			currenTime := time.Now()
			date1 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day())
			date2 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day()+1)
		}

		income, err := repository.GetIncome(r.Context(), params["id"], date1, date2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(income)
	}
}

func GetAllTransaction(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		date1 := r.URL.Query().Get("date1")
		date2 := r.URL.Query().Get("date2")

		token, err := helpers.DecodeJwt(w, r, s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*models.AppClaims)

		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return

		}

		if date1 == "" || date2 == "" {
			currenTime := time.Now()
			date1 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day())
			date2 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day()+1)
		}

		income, err := repository.GetAllTransaction(r.Context(), claims.UserId, date1, date2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(income)
	}
}
func DeleteIncome(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := repository.DeleteIncome(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.DeleteUPdateReponse{
			Message: "Income Delete",
		})
	}
}

// func IncomeAndBill(s server.Server) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)

// 		// date1 := r.URL.Query().Get("date1")
// 		// date2 := r.URL.Query().Get("date2")

// 		// if date1 == "" || date2 == "" {
// 		// 	currenTime := time.Now()
// 		// 	date1 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day())
// 		// 	date2 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day()+1)
// 		// }

// 		// income, err := repository.GetIncome(r.Context(), params["id"], date1, date2)
// 		// if err != nil {
// 		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 		// 	return
// 		// }

// 		// bill, err := repository.GetBills(r.Context(), params["id"], date1, date2)
// 		// if err != nil {
// 		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 		// 	return
// 		// }

// 		incomeAndBill, err := repository.GetIncomeAndBill(r.Context(), params["id"])
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(incomeAndBill)
// 	}
// }
