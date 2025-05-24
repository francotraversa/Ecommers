package main

import (
	"log"

	"github.com/francotraversa/GOLANG/routes"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()
	db.Init()
	e := echo.New()
	routes.RegisterPublicsRoutes(e)
	routes.RegisterPrivatesRoutes(e)

	log.Println("Servidor corriendo en :8080")
	e.Logger.Fatal(e.Start(":8080"))

}
