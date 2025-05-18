package models

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"contrasena"` // hash
	Email    string `json:"email"`
	Master   bool   `json:"master"`
}
