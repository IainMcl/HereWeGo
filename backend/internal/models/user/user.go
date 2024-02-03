package models

import (
	"errors"

	"github.com/IainMcl/HereWeGo/internal/database"
	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/services"
)

type UserRole int

const (
	Admin UserRole = iota
	StandardUser
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" db:"role"`
}

func (u *User) CreateUser() error {
	p, err := services.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = p
	if !services.ValidateEmail(u.Email) {
		return errors.New("invalid email")
	}
	if err := u.addUserToDb(); err != nil {
		return err
	}
	return nil
}

func (u *User) addUserToDb() error {
	s := database.New().Db()
	tx := s.MustBegin()
	tx.MustExec(`
		INSERT INTO users 
			(username, email, password) 
			VALUES ($1, $2, $3)
			`, u.Username, u.Email, u.Password)

	tx.Commit()
	return nil
}

func AuthenticateUser(email, password string) (User, error) {
	var u User
	s := database.New().Db()
	err := s.Get(&u, `
		SELECT 
			id, username, email, password, role
		FROM users 
		WHERE email = $1
		`, email)
	if err != nil {
		logging.Warn("Error getting user from db: ", err)
		return u, err
	}
	if !services.CheckPasswordHash(password, u.Password) {
		return u, errors.New("invalid password")
	}
	return u, nil
}
