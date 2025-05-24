package models

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // hash
	Email    string `json:"email"`
	Master   bool   `json:"master"`
}
type UsersLogin struct {
	Username string `json:"username"`
	Password string `json:"contrasena"`
}
