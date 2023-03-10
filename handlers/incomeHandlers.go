package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/osmait/admin-finanzas/helpers"
	"github.com/osmait/admin-finanzas/models"
	"github.com/osmait/admin-finanzas/repository"
	"github.com/osmait/admin-finanzas/server"
	"github.com/segmentio/ksuid"
)

func InsertTransaction(s server.Server) http.HandlerFunc {
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
		var transaction = models.Transaction{}
		err = json.NewDecoder(r.Body).Decode(&transaction)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		validate := validator.New()
		err = validate.Struct(transaction)
		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		transaction.Id = id.String()
		transaction.UserId = claims.UserId

		err = repository.InsertTransaction(r.Context(), &transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)

	}
}
func GetTransaction(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		date1 := r.URL.Query().Get("date1")
		date2 := r.URL.Query().Get("date2")

		if date1 == "" || date2 == "" {
			currenTime := time.Now()
			date1 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day())
			date2 = fmt.Sprintf("%d/%d/%d", currenTime.Year(), currenTime.Month(), currenTime.Day()+1)
		}

		transaction, err := repository.GetTransaction(r.Context(), params["id"], date1, date2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)
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

		transaction, err := repository.GetAllTransaction(r.Context(), claims.UserId, date1, date2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)
	}
}

func UpdateTransaction(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
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
		var transaction = models.Transaction{}
		err = json.NewDecoder(r.Body).Decode(&transaction)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		transaction.UserId = claims.UserId
		transaction.Id = params["id"]
		err = repository.UpdateTransaction(r.Context(), params["id"], &transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)

	}
}

func DeleteTransaction(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := repository.DeleteTransaction(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.DeleteUPdateReponse{
			Message: "Transaction Delete",
		})
	}
}

func GetBalace(s server.Server) http.HandlerFunc {
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
		total, err := repository.GetBalance(r.Context(), claims.UserId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(total)

	}
}
