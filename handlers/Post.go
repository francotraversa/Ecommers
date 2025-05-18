package handlers

import (
	"net/http"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
)

func PostProduct(c echo.Context) error {
	var producto models.Producto

	if err := c.Bind(&producto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	identiti := db.ExistsTaskID(producto.ID, c)
	if identiti == -1 {
		_, err := db.DB.Exec(
			"INSERT INTO productos (id, titulo, description, precio, stock) VALUES ($1, $2, $3, $4, $5)",
			producto.ID, producto.Title, producto.Description, producto.Price, producto.Stock)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando la tarea"})
		}
	} else {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "El ID ya existe"})
	}
	return c.JSON(http.StatusCreated, producto)
}
