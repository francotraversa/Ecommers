package handlers

import (
	"database/sql"
	"net/http"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	var productos []models.Producto

	rows, err := db.DB.Query("SELECT id, titulo, descripcion, precio, stock FROM productos order by titulo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al consultar la base de datos"})
	}
	defer rows.Close()

	for rows.Next() {
		var producto models.Producto
		if err := rows.Scan(&producto.ID, &producto.Title, &producto.Description, &producto.Price, &producto.Stock); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al leer las tareas"})
		}
		productos = append(productos, producto)
	}

	if len(productos) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "No hay tareas registradas"})
	}

	return c.JSON(http.StatusOK, productos)
}

func GetTaskByID(c echo.Context) error {
	id := c.Param("id")

	var productoyid models.Producto

	err := db.DB.QueryRow("SELECT id, titulo, descripcion, precio, stock FROM productos WHERE id = $1", id).
		Scan(&productoyid.ID, &productoyid.Title, &productoyid.Description, &productoyid.Price, &productoyid.Stock)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Tarea no encontrada"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
	}

	return c.JSON(http.StatusOK, productoyid)
}
