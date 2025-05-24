package db

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/francotraversa/GOLANG/models"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error abriendo la conexión:", err)
	}

	if err != nil {
		log.Fatal("No se pudo conectar a la DB:", err)
	}

	log.Println("Base de datos conectada correctamente.")

}

func ExistsUsuarioID(id int, c echo.Context) int {
	var user models.Users
	result := DB.Table("usuarios").First(&user, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return -1
		}
		c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
		return -1
	}
	return user.ID
}
func ExistsProductoID(id int, c echo.Context) int {
	var producto models.Producto
	result := DB.Table("productos").First(&producto, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return -1
		}
		c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
		return -2
	}
	return producto.ID
}
func GetProductbyID(id int, c echo.Context) (error, bool) {
	var producto models.AskProducto
	if err := DB.Table("productos").Where("id = ?", id).First(&producto).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "El Producto ya existe"}), true
	}
	return c.JSON(http.StatusOK, producto), false

}
