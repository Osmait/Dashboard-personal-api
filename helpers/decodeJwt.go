package helpers

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/osmait/admin-finanzas/models"
	"github.com/osmait/admin-finanzas/server"
)

func DecodeJwt(w http.ResponseWriter, r *http.Request, s server.Server) (*jwt.Token, error) {
	tokeString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokeString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}
