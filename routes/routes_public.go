package routes

import (
	"github.com/francotraversa/GOLANG/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterPublicsRoutes(e *echo.Echo) {
	e.GET("/productos", handlers.GetProducts)
	e.GET("/productos/", handlers.GetProductporID)
	e.POST("/login", handlers.LoginUser)
}
