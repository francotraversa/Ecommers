package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProducts(c echo.Context) error {
	var productos []models.Producto

	if err := db.DB.Table("productos").Find(&productos).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al consultar la base de datos"})
	}

	if len(productos) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "No hay tareas registradas"})
	}

	return c.JSON(http.StatusOK, productos)
}

func GetProductporID(c echo.Context) error {
	var productobyid models.AskProducto
	var producto models.Producto
	if err := c.Bind(&productobyid); err != nil {
		fmt.Println("Bind error:", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	if err := db.DB.Table("productos").Where("id = ?", productobyid.ID).First(&producto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Producto no encontrado"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al consultar la base de datos"})
	}
	return c.JSON(http.StatusOK, producto)

}
