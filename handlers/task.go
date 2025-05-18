package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/francotraversa/GOLANG/models"
	db "github.com/francotraversa/GOLANG/storage"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func PostTask(c echo.Context) error {
	var newTask models.Task

	if err := c.Bind(&newTask); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	identiti := db.ExistsTaskID(newTask.ID, c)
	if identiti == -1 {
		_, err := db.DB.Exec(
			"INSERT INTO tasks (id, title, description, completed) VALUES ($1, $2, $3, $4)",
			newTask.ID, newTask.Title, newTask.Description, newTask.Completed,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error insertando la tarea"})
		}
	} else {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "El ID ya existe"})
	}
	return c.JSON(http.StatusCreated, newTask)
}

func GetTasks(c echo.Context) error {
	var tasks []models.Task

	rows, err := db.DB.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al consultar la base de datos"})
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al leer las tareas"})
		}
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "No hay tareas registradas"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c echo.Context) error {
	id := c.Param("id")

	var task models.Task

	err := db.DB.QueryRow("SELECT id, title, description, completed FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Tarea no encontrada"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error consultando la base de datos"})
	}

	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	identiti := db.ExistsTaskID(id, c)
	if identiti == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "El ID no existe"})
	}

	var updatedTask models.Task

	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al parsear JSON"})
	}

	_, err = db.DB.Exec(
		"UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4",
		updatedTask.Title, updatedTask.Description, updatedTask.Completed, identiti,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error actualizando la tarea"})
	}

	updatedTask.ID = identiti
	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	identiti := db.ExistsTaskID(id, c)
	if identiti == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "El ID no existe"})
	}

	_, err = db.DB.Exec("DELETE FROM tasks WHERE id = $1", identiti)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error eliminando la tarea"})
	}

	return c.NoContent(http.StatusNoContent)
}
func DeleteAllTasks(c echo.Context) error {
	_, err := db.DB.Exec("DELETE FROM tasks")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error eliminando las tareas"})
	}

	return c.NoContent(http.StatusNoContent)
}

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
	if user.Username == "admin" && user.Password == "traversa" {
		log.Println("Usuario admin")
		sess, _ := session.Get("session", c)
		sess.Values["username"] = user.Username
		sess.Save(c.Request(), c.Response())
		return c.JSON(http.StatusOK, echo.Map{"message": "Se logeo correctamente como admin"})
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
