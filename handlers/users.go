package handlers

import (
	"log"
	"net/http"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/google/uuid"
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

	result := db.DB.Table("usuarios").Create(&newUser)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando el usuario"})
	}

	useregistered := GetUserByName(newUser.Username)

	if useregistered == nil {
		log.Println("Usuario no encontrado: ", newUser.Username)
	}
	result = db.DB.Table("acciones_usuario").Create(&models.Accion{
		//ID_Usuario:  useregistered.ID, Aca tendri que ser el ID del usuario que lo creo
		ID_Articulo: newUser.ID,
		Tipo_accion: "Alta Usuario"})

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando el usuario en Acciones_Usuario"})
	}

	return c.JSON(http.StatusCreated, &useregistered)
}

func LoginUser(c echo.Context) error {
	var userrequest models.UsersLogin
	if err := c.Bind(&userrequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	user := GetUserByName(userrequest.Username)
	if user == nil {
		log.Println("Usuario no encontrado: ", userrequest.Username)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Usuario no encontrado"})
	}

	log.Print("Usuario encontrado: ", userrequest.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userrequest.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Contraseña incorrecta"})
	}

	log.Println("Se inicio sesión correctamente por el usuario: ", user.Username)

	uuid := uuid.New().String()
	err := db.DB.Table("sesiones").Create(&models.Sesion{
		Usuario_id: user.ID,
		Uuid:       uuid}).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando la sesión"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Inicio de sesión exitoso",
		"UUID":    uuid})
}

func Logout(c echo.Context) error {
	//err := db.DB.Delete("sesiones").Create(&models.Sesion{
	//	Usuario_id: user.ID,
	//	Uuid:       uuid}).Error

	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando la sesión"})
	//}
	return c.JSON(http.StatusOK, echo.Map{"message": "Sesión cerrada"})
}

func DisableUser(c echo.Context) error {
	var userrequest models.UsersLogin
	if err := c.Bind(&userrequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}
	disableuser := GetUserByName(userrequest.Username)
	err := db.DB.Delete(&models.Users{}, disableuser.ID).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error eliminando el usuario"})
	}
	//Insert := db.DB.Table("acciones_usuario").Create(&models.Accion{
	//	ID_Usuario: 213123, //
	//	ID_Articulo: disableuser.ID,
	//	Tipo_accion: "Alta Usuario"})

	//if Insert.Error != nil {
	//	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando la tarea"})
	//}

	// Aca tendriamos que poner una columna en usuarios que se llame deleted y que si esta en null no se elimino y si tiene una fecha se elimino
	return c.NoContent(http.StatusNoContent)
}

func GetUser(ID int) *models.Users {
	var user models.Users
	db.DB.Table("usuarios").Where("id = ?", ID).First(&user)
	return &user
}
func GetUserByName(username string) *models.Users {
	var user models.Users
	db.DB.Table("usuarios").Where("username = ?", username).First(&user)
	return &user
}
