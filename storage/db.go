package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var DB *sql.DB

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
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error abriendo la conexión:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la DB:", err)
	}

	log.Println("Base de datos conectada correctamente.")

}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error cerrando la conexión a la base de datos:", err)
	}
	log.Println("Conexión a la base de datos cerrada.")
}

func ExistsTaskID(id int, c echo.Context) int {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
	}
	if !exists {
		return -1
	}
	return id
}
