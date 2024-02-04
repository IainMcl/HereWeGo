package models

import (
	"errors"

	"github.com/IainMcl/HereWeGo/internal/database"
	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/services"
	"github.com/jmoiron/sqlx"
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

func (u *User) CreateUser(db *sqlx.DB) error {
	p, err := services.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = p
	if !services.ValidateEmail(u.Email) {
		return errors.New("invalid email")
	}
	if err := u.addUserToDb(db); err != nil {
		return err
	}
	return nil
}

func (u *User) addUserToDb(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec(`
		INSERT INTO users 
			(username, email, password) 
			VALUES ($1, $2, $3)
			`, u.Username, u.Email, u.Password)

	tx.Commit()
	return nil
}

func (u *User) UpdatePassword(db *sqlx.DB, password string) error {
	p, err := services.HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = p
	err = u.UpdateUser(db)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec(`
		UPDATE users 
		SET username = $1, email = $2, password = $3
		WHERE id = $4
		`, u.Username, u.Email, u.Password, u.ID)
	tx.Commit()
	return nil
}

func AuthenticateUser(db *sqlx.DB, email, password string) (User, error) {
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

func GetUserByEmail(db *sqlx.DB, email string) (User, error) {
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
	return u, nil
}
