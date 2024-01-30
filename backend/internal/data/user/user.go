package data

import (
	"github.com/IainMcl/HereWeGo/internal/database"
	models "github.com/IainMcl/HereWeGo/internal/models/user"
)

func CreateUser(user models.User) error {
	s := database.New().Db()
	tx := s.MustBegin()
	tx.MustExec(`
		INSERT INTO users 
			(username, email, password) 
			VALUES ($1, $2, $3)
			`, user.UserName, user.Email, user.Password)

	tx.Commit()
	return nil
}
