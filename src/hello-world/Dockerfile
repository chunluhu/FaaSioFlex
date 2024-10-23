# Usar una imagen base de Go
FROM golang:1.23.2 AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos Go al contenedor
COPY go.mod . 
COPY go.sum . 

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o hello-world .

# Usar una imagen más ligera para ejecutar la aplicación
FROM alpine:latest

# Copiar el binario compilado desde la imagen de construcción
COPY --from=builder /app/hello-world .

# Exponer el puerto que la aplicación usará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./hello-world"]
