package main

import (
	"log"
	"net/http"

	"github.com/francotraversa/GOLANG/handlers"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()
	db.Init()
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret-key"))))
	e.GET("/tasks", handlers.GetTasks)
	e.GET("/tasks/:id", handlers.GetTaskByID)
	e.POST("/login", handlers.LoginUser)

	//////////////////////////////////////////////////////////

	//e.POST("/tasks", handlers.PostTask) PUBLICA
	//e.PUT("/tasks/:id", handlers.UpdateTask) PUBLICA
	//e.DELETE("/tasks/:id", handlers.DeleteTask) PUBLICA
	//e.DELETE("/delete/all", handlers.DeleteTask) PUBLICA

	////////////////////////////////////////////////////////////

	r := e.Group("/tasks")
	r.Use(AuthMiddleware)

	r.POST("/tasks", handlers.PostTask)
	r.PUT("/tasks", handlers.UpdateTask)
	r.DELETE("/:id", handlers.DeleteTask)
	r.DELETE("/delete/all", handlers.DeleteAllTasks)
	r.POST("/users/", handlers.RegisterUser)
	r.POST("/close", handlers.Logout)
	//r.GET("/users/:id", handlers.GetUserByID)
	//r.PUT("/users/:id", handlers.UpdateUser)

	// Inicia el servidor
	log.Println("Servidor corriendo en :8080")
	e.Logger.Fatal(e.Start(":8080"))

}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		username := sess.Values["username"]
		if username == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "No autorizado"})
		}
		return next(c)
	}
}
