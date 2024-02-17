package services

import (
	"net/mail"

	mailService "github.com/IainMcl/HereWeGo/internal/mail"
	"github.com/IainMcl/HereWeGo/internal/settings"
	"github.com/IainMcl/HereWeGo/internal/util"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func SendResetPasswordEmail(email, username string, userId int64, token util.OTPToken) error {
	mailService := mailService.New()
	expiresAt := token.Expiration.Format(settings.AppSettings.UserDateTimeFormat)
	data := map[string]interface{}{
		"name":        username,
		"token":       token.Token,
		"userID":      userId,
		"frontendURL": "https://google.com", // TODO: Change this to the frontend URL
		"expiration":  token.ExpirationMinutes,
		"exact":       expiresAt,
	}
	err := mailService.SendMailTemplate(email, "password_reset.tmpl", data)
	if err != nil {
		return err
	}
	return nil
}

func ObscureEmail(email string) string {
	atIndex := 0
	for i, c := range email {
		if string(c) == "@" {
			atIndex = i
			break
		}
	}
	return email[:1] + "..." + email[atIndex-1:]
}
