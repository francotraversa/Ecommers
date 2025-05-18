package handlers

import (
	"net/http"
	"strconv"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
)

func UpdateProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	identiti := db.ExistsTaskID(id, c)
	if identiti == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "El ID no existe"})
	}

	var updatedproduct models.Producto

	if err := c.Bind(&updatedproduct); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	_, err = db.DB.Exec(
		"UPDATE productos SET title = $1, description = $2, precio = $3, stock = $4 WHERE id = $5",
		updatedproduct.Title, updatedproduct.Description, updatedproduct.Price, updatedproduct.Stock, identiti,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error actualizando la tarea"})
	}

	updatedproduct.ID = identiti
	return c.JSON(http.StatusOK, updatedproduct)
}
