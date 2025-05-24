package routes

import (
	"github.com/francotraversa/GOLANG/handlers"
	"github.com/francotraversa/GOLANG/middleware"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func RegisterPrivatesRoutes(e *echo.Echo) {
	r := e.Group("/loged")
	r.Use(middleware.AuthMiddleware)
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret-key"))))
	r.POST("/post", handlers.PostProduct)
	r.PUT("/update", handlers.UpdateProduct)
	r.DELETE("/delete", handlers.DeleteProduct)
	r.POST("/register", handlers.RegisterUser)
	r.POST("/close", handlers.Logout)
}
