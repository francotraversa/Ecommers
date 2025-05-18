package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"contrasena"` // hash
	Email    string `json:"email"`
	Master   bool   `json:"master"`
}

/*
login

"username": "admin",
"contrasena": "admin",



consultas
{
  "id": "21",
  "title": "travesrsa",
  "description" : "fran",
  "completed" : "true"
}

*/
