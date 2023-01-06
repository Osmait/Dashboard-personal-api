package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/osmait/admin-finanzas/helpers"
	"github.com/osmait/admin-finanzas/models"
	"github.com/osmait/admin-finanzas/repository"
	"github.com/osmait/admin-finanzas/server"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_COST = 10
)

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user = models.User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println(user)
		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), HASH_COST)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := ksuid.NewRandom()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tokenAuth, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Id = id.String()
		user.Password = string(hashedPassword)
		user.Token = tokenAuth.String()

		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Password = ""
		go helpers.SendMail(tokenAuth.String(), user.Email)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func GetUserById(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		user, err := repository.GetUserById(r.Context(), params["id"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func UserTokenAuth(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		user, err := repository.GetUserByToken(r.Context(), params["token"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if user.Token == "" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user.Confirmed = true
		user.Token = ""
		err = repository.UpdateConfirmed(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.DeleteUPdateReponse{
			Message: "User Confirmed",
		})

	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = models.SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := repository.GetUserEmail(r.Context(), request.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if !user.Confirmed {
			http.Error(w, "User No Autenticado", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)

			return
		}
		token, err := helpers.SignToken(user.Id, s)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.LoginResponse{
			Token: token,
			User:  user,
		})

	}
}

func Perfil(s server.Server) http.HandlerFunc {
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(claims)
	}

}
