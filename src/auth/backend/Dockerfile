 # Usa una imagen de Go como base
 FROM golang:1.18-alpine AS builder

 # Establece el directorio de trabajo
 WORKDIR /app

 # Copia el archivo go.mod y go.sum
 COPY go.mod ./
 COPY go.sum ./

 # Descarga los módulos Go
 RUN go mod download

 # Copia el código fuente del backend en el contenedor
 COPY . .

 # Compila la aplicación de Go
 RUN go build -o /auth

 # Etapa para crear el contenedor final
 FROM alpine:latest

 # Crea directorio para la configuración de DEx
 RUN mkdir -p /etc/dex

 # Copia el binario del backend y el archivo de configuración de DEx
 COPY --from=builder /auth /auth
 COPY backend/dex/config.yaml /etc/dex/config.yaml

 # Instala DEx
 RUN apk add --no-cache curl && \
     curl -L https://github.com/dexidp/dex/releases/download/v2.29.0/dex_v2.29.0_linux_amd64.tar.gz | tar xz && \
     mv dex /usr/local/bin/

 # Expone los puertos necesarios
 EXPOSE 5555 5556

 # Comando para correr la aplicación y DEx
 CMD ["/bin/sh", "-c", "/auth & dex serve /etc/dex/config.yaml"]
