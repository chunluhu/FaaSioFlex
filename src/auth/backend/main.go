package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "your_project_name/config"
    "your_project_name/routes"
)

func main() {
    // Cargar configuración
    config.Load()

    // Crear un nuevo router
    router := mux.NewRouter()

    // Configurar las rutas
    routes.SetupAuthRoutes(router)
    routes.SetupFunctionRoutes(router)

    // Iniciar el servidor
    log.Println("Servidor escuchando en :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
