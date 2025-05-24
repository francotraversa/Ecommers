package models

type Accion struct {
	ID_Usuario  int    `json:"id_usuario"`
	ID_Articulo int    `json:"id_articulo"`
	Tipo_accion string `json:"Tipo_accion"`
}
