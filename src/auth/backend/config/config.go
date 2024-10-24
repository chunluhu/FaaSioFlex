package config

import (
    "log"
    "os"
)

func Load() {
    // Cargar variables de entorno y configuraciones
    // Aquí puedes cargar la base de datos y otros parámetros
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error al cargar .env file")
    }
}
