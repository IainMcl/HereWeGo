package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/settings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var JwtSecret []byte

type Claims struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserId int64  `json:"user_id"`
	jwt.RegisteredClaims
}

// func (c Claims) Valid() error {
// 	// TODO: Add more validation
// 	return nil
// }

func GenerateToken(username, email string, userId int64) (string, error) {
	// Set custom claims
	claims := &Claims{
		username,
		email,
		userId,
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

// Using an in-memory blacklist for simplicity
var jwtBalcklist []string

func AddBlacklist(jti string) {
	logging.Info("Adding token to blacklist")
	jwtBalcklist = append(jwtBalcklist, jti)
	go cleanBlacklist()
}

func IsBlacklisted(jti string) bool {
	for _, v := range jwtBalcklist {
		if v == jti {
			logging.Debug("Attempted use of blacklisted token", jti)
			return true
		}
	}
	return false
}

func cleanBlacklist() {
	logging.Debug("Cleaning blacklist")
	now := time.Now()
	for i, v := range jwtBalcklist {
		exp, err := getExpirationTime(v)
		if err != nil {
			continue
		}
		if now.After(exp.Time) {
			jwtBalcklist = append(jwtBalcklist[:i], jwtBalcklist[i+1:]...)
		}
	}
}

func getExpirationTime(tokenString string) (jwt.NumericDate, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &Claims{})
	if err != nil {
		return jwt.NumericDate{}, err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return *claims.ExpiresAt, nil
	}

	return jwt.NumericDate{}, fmt.Errorf("exp claim not found")
}

func GetEmailFromToken(tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &Claims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims.Email, nil
	}

	return "", fmt.Errorf("email claim not found")
}

// Get the claims from a token
//
// tokenString - either a jwt token or a string with the format "Bearer <jwt token>"
func ClaimsFromToken(tokenString string) (Claims, error) {
	if tokenString == "" {
		logging.Warn("Empty token")
		return Claims{}, errors.New("invalid token")
	}
	if strings.HasPrefix(tokenString, "Bearer ") && len(tokenString) > 7 {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &Claims{})
	if err != nil {
		return Claims{}, err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		logging.Debug("Claims parsed for: ", claims.UserId)
		return *claims, nil
	}

	return Claims{}, fmt.Errorf("error parsing claims")
}

func GetUserId(c echo.Context) (int64, error) {
	user, err := ClaimsFromToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		return 0, err
	}
	return user.UserId, nil
}
