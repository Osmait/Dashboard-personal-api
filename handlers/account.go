package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/osmait/admin-finanzas/helpers"
	"github.com/osmait/admin-finanzas/models"
	"github.com/osmait/admin-finanzas/repository"
	"github.com/osmait/admin-finanzas/server"
	"github.com/segmentio/ksuid"
)

func InsertAccount(s server.Server) http.HandlerFunc {
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

		var account = models.Account{}
		err = json.NewDecoder(r.Body).Decode(&account)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(account)
		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		account.Id = id.String()

		account.User_id = claims.UserId
		err = repository.InsertAccount(r.Context(), &account)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(account)
	}
}

func GetAccounts(s server.Server) http.HandlerFunc {
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
		accounts, err := repository.GetAcoounts(r.Context(), claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accounts)
	}
}

func DeleteAccount(s server.Server) http.HandlerFunc {
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
		err = repository.DeleteAccount(r.Context(), params["id"], claims.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.DeleteUPdateReponse{
			Message: "Account Delete",
		})
	}
}
