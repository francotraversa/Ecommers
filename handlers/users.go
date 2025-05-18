package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	var newUser models.Users

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error encriptando contraseña"})
	}

	newUser.Password = string(hashedPassword)

	_, err = db.DB.Exec("INSERT INTO usuarios (username, contrasena) VALUES ($1, $2)",
		newUser.Username, string(hashedPassword))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando el usuario"})
	}
	err = db.DB.QueryRow("SELECT id FROM usuarios WHERE username = $1", newUser.Username).Scan(&newUser.ID)
	if err == sql.ErrNoRows {
		log.Println("Usuario no encontrado: ", newUser.Username)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
	}
	//newUser.Password = ""
	return c.JSON(http.StatusOK, echo.Map{"Message": "Usuario agregado", "username": newUser.Username, "id": newUser.ID, "password": newUser.Password})
}

func LoginUser(c echo.Context) error {

	var user models.Users
	var hashedpassword string
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	err := db.DB.QueryRow("SELECT username, contrasena FROM usuarios WHERE username = $1",
		user.Username).Scan(&user.Username, &hashedpassword)

	if err == sql.ErrNoRows {
		log.Println("Usuario no encontrado: ", user.Username)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "usuario no encontrado"})

	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Contraseña incorrecta"})
	}

	log.Println("Se inicio sesión correctamente por el usuario: ", user.Username)

	sess, _ := session.Get("session", c)
	sess.Values["username"] = user.Username
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "Inicio de sesión exitoso",
		"username": user.Username})
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1 // expira ya
	sess.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, echo.Map{"message": "Sesión cerrada"})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM usuarios WHERE id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error eliminando el usuario"})
	}
	return c.NoContent(http.StatusNoContent)
}
