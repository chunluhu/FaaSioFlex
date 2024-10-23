package main

import (
    "fmt"
    "net/http"  
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hola Mundo - tres veces  - ")
}

func main() {
    http.HandleFunc("/", helloWorld) // Configura la ruta
    fmt.Println("Servidor escuchando en http://localhost:8080")
    http.ListenAndServe(":8080", nil) // Inicia el servidor
}
