package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/osmait/admin-finanzas/models"
	"github.com/osmait/admin-finanzas/server"
)

func SignToken(id string, s server.Server) (string, error) {
	claims := models.AppClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
