# Usamos una imagen base de Go
FROM golang:1.23

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos del proyecto al contenedor
COPY . .

# Compilamos la aplicación
RUN go build -o funciones_api

# Exponemos el puerto 5000
EXPOSE 5000

# Comando para ejecutar la aplicación
CMD ["./funciones_api"]
