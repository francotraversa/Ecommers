package models

type Producto struct {
	ID          int     `json:"id"`
	Title       string  `json:"titulo"`
	Description string  `json:"descripcion"`
	Price       float32 `json:"precio"`
	Stock       int     `json:"stock"`
}

type Venta struct {
	ID          int     `json:"id"`
	Title       string  `json:"titulo"`
	Description string  `json:"descripcion"`
	Price       float32 `json:"precio"`
	Cantidad    int     `json:"cantidad"`
	Total       float32 `json:"total"`
	Fecha       string  `json:"fecha"`
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
