# Gestor de Tareas con Go y Echo 🚀

Desarrollado por **Franco Traversa**

Este proyecto es una API RESTful construida con [Go](https://golang.org/) y el framework [Echo](https://echo.labstack.com/) para gestionar tareas. Implementa autenticación de usuarios utilizando sesiones, y permite que usuarios anónimos consulten tareas, mientras que los registrados pueden crear, editar y eliminar.

## 🧰 Tecnologías utilizadas

- Go 1.21+
- Echo framework
- PostgreSQL
- Middleware JWT (opcional)
- Sesiones con `gorilla/sessions`
- Bcrypt para encriptar contraseñas

## 📦 Instalación

1. Cloná este repositorio:

```bash
git clone https://github.com/tuusuario/nombre-del-repo.git
cd nombre-del-repo

Configurá tu base de datos PostgreSQL:

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    contrasena TEXT NOT NULL,
    email TEXT,
    master BOOLEAN DEFAULT false
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    titulo TEXT NOT NULL,
    descripcion TEXT,
    completada BOOLEAN DEFAULT false,
    usuario_id INTEGER REFERENCES usuarios(id)
);

Rutas Públicas
POST /register → Registrar un nuevo usuario

POST /login → Iniciar sesión

GET /tasks → Listar tareas públicas

GET /tasks/:id → Ver una tarea específica

Rutas Protegidas (requieren sesión iniciada)
POST /tasks → Crear tarea

PUT /tasks/:id → Editar tarea

DELETE /tasks/:id → Eliminar tarea

GET /logout → Cerrar sesión

⚙️ Funcionalidades destacadas
✔️ CRUD de tareas

🔐 Autenticación por sesión

🔑 Contraseñas encriptadas con bcrypt

🧪 Control de acceso simple por usuario

👤 Admin validado directamente (sin pasar por base de datos)
