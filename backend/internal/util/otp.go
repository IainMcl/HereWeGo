package util

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/settings"
	"github.com/jmoiron/sqlx"
)

type OTPToken struct {
	Token             string `db:"code"`
	ExpirationMinutes int
	Expiration        time.Time `db:"expires_at"`
	IssuedAt          time.Time `db:"created_at"`
	UserId            int64     `db:"user_id"`
}

func NewOTPToken(db *sqlx.DB, userId int64) (OTPToken, error) {
	tokenString := generateOTP()

	now := time.Now()
	token := OTPToken{
		Token:             tokenString,
		ExpirationMinutes: settings.OTPSettings.ExpirationMinutes,
		Expiration:        now.Add(time.Duration(settings.OTPSettings.ExpirationMinutes) * time.Minute),
		IssuedAt:          now,
		UserId:            userId,
	}
	logging.Info("Creating new OTP token for user " + strconv.Itoa(int(token.UserId)))
	err := token.SetToken(db)
	if err != nil {
		return OTPToken{}, err
	}
	return token, nil
}

func (t OTPToken) SetToken(db *sqlx.DB) error {
	_, err := db.Exec(
		"INSERT INTO otp (user_id, code, created_at, expires_at) VALUES ($1, $2, $3, $4)",
		t.UserId, t.Token, t.IssuedAt, t.Expiration)
	return err
}

func (t OTPToken) IsExpired() bool {
	return time.Now().After(t.Expiration)
}

func (t OTPToken) IsValid(token, email string) bool {
	return false

}

func generateOTP() string {
	code := make([]rune, settings.OTPSettings.Length)
	for i := 0; i < settings.OTPSettings.Length; i++ {
		code[i] = rune(48 + rand.Intn(10))
	}
	return string(code)
}

func UseToken(db *sqlx.DB, token string, userId int64) error {
	_, err := db.Exec("DELETE FROM otp WHERE code = $1 and user_id = $2", token, userId)
	return err
}

func GetTokenFromDB(db *sqlx.DB, token, email string) (OTPToken, error) {
	var t OTPToken
	err := db.Get(&t,
		`SELECT 
			o.user_id
			, code
			, o.created_at
			, expires_at 
		FROM otp o 
		left join users u 
		on o.user_id = u.id 
		WHERE code = $1 AND email = $2
		order by o.created_at desc
		LIMIT 1`,
		token, email)
	return t, err
}
