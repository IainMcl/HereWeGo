package util

import (
	"time"

	"github.com/IainMcl/HereWeGo/internal/settings"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret []byte

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// func (c Claims) Valid() error {
// 	// TODO: Add more validation
// 	return nil
// }

func GenerateToken(username, email string) (string, error) {
	// Set custom claims
	claims := &Claims{
		username,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(settings.AppSettings.JwtDurationHours))),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}
