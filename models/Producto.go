package models

type Producto struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Title       string  `json:"titulo" gorm:"column:titulo"`
	Description string  `json:"descripcion" gorm:"column:descripcion"`
	Price       float64 `json:"precio" gorm:"column:precio"`
	Stock       int     `json:"stock" gorm:"column:stock"`
}

type Venta struct {
	ID       int `json:"id" gorm:"primaryKey"`
	Cantidad int `json:"cantidad"`
}

type AskProducto struct {
	ID int `json:"id"`
}

type AksVenta struct {
	ArticuloID int `json:"id"`
}

type UpdateProduct struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Stock int     `json:"stock" gorm:"column:stock"`
	Price float64 `json:"precio" gorm:"column:precio"`
}
