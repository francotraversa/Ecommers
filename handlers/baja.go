package handlers

import (
	"net/http"
	"time"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
)

func DeleteProduct(c echo.Context) error {
	var deletedproduct models.AskProducto
	if err := c.Bind(&deletedproduct); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}
	exist := db.ExistsProductoID(deletedproduct.ID, c)

	switch exist {
	case -1:
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Producto no encontrado"})
	case -2:
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
	}
	err := db.DB.Model(&models.Producto{}).Where("id = ?", exist).Update("unsubscribe", time.Now()).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error eliminando el producto"})
	}

	_ = db.DB.Model(&models.Accion{}).Create(&models.Accion{
		ID_Usuario:  exist,
		ID_Articulo: deletedproduct.ID,
		Tipo_accion: "delete",
	}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error registrando la acción"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Producto eliminado"})
}
