package handlers

import (
	"net/http"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
)

func UpdateProduct(c echo.Context) error {
	var UpdateProduct models.UpdateProduct
	if err := c.Bind(&UpdateProduct); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	producto := db.ExistsProductoID(UpdateProduct.ID, c)
	switch producto {
	case -1:
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Producto no encontrado"})
	case -2:
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al consultar la base de datos"})
	}

	if err := db.DB.Model(&models.Producto{}).Where("id = ?", UpdateProduct.ID).Updates(models.Producto{
		Price: UpdateProduct.Price,
		Stock: UpdateProduct.Stock,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al actualizar el producto"})
	}
	return c.JSON(http.StatusOK, producto)
}
