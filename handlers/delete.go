package handlers

import (
	"net/http"
	"strconv"

	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo/v4"
)

func DeleteProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	identiti := db.ExistsTaskID(id, c)
	if identiti == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "El ID no existe"})
	}

	_, err = db.DB.Exec("DELETE FROM producto WHERE id = $1", identiti)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error eliminando la tarea"})
	}

	return c.NoContent(http.StatusNoContent)
}
func DeleteAllTasks(c echo.Context) error {
	_, err := db.DB.Exec("drop table productos")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error eliminando las tareas"})
	}

	return c.NoContent(http.StatusNoContent)
}
