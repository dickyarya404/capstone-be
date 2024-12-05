package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sawalreverr/recything/config"
)

type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userID string, role string) (string, error) {
	claims := &JwtCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.GetConfig().Server.JWTSecret
	return token.SignedString([]byte(secret))
}
