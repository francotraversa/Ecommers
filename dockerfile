# Usa la imagen oficial de Go como base
FROM golang:1.23.3-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Descarga las dependencias
RUN go mod download

# Compila el binario
RUN go build -o main ./cmd/main.go

# Expone el puerto donde corre tu API
EXPOSE 8080

# Comando para ejecutar la API
CMD ["./main"]
