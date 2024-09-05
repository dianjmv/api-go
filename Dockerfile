# Usa una imagen base oficial de Go
FROM golang:1.16-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente de la aplicación
COPY . .

# Compila la aplicación
RUN go build -o main ./cmd

# Expone el puerto en el que la app correrá
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
