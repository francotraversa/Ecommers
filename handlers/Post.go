package handlers

import (
	"net/http"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
)

func PostProduct(c echo.Context) error {
	//if user == nil {
	//	return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Usuario no autenticado"})
	//}
	//if user.Master { todo esto }
	var producto models.Producto

	if err := c.Bind(&producto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}
	exist := db.ExistsProductoID(producto.ID, c)

	switch exist {
	case -1:
		result := db.DB.Table("productos").Create(&producto)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando el producto"})
		}
	case -2:
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error: En la consuta de la base de datos"})
	}
	return c.JSON(http.StatusOK, echo.Map{"Existe": "Existe: El producto ya existe"})
}
