package models

type UserRole int

const (
	Admin UserRole = iota
	StandardUser
)

type User struct {
	ID       int64    `json:"id"`
	UserName string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}
