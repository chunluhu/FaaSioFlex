package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Almacenamos las funciones registradas en un mapa
var funcionesRegistradas = make(map[string]func() interface{})
var mu sync.Mutex

// Estructura para la solicitud de registro
type Registro struct {
	Nombre string `json:"nombre"`
	Codigo string `json:"codigo"`  
}

// Función para registrar una función
func registrarFuncion(w http.ResponseWriter, r *http.Request) {
	var reg Registro
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	mu.Lock()
	defer mu.Unlock()

	// Registramos la función
	funcionesRegistradas[reg.Nombre] = func() interface{} {
		var resultado interface{}
		// Usamos el paquete fmt para simular la ejecución del código
		fmt.Sscanf(reg.Codigo, "%v", &resultado)
		return resultado
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": fmt.Sprintf("Función %s registrada exitosamente", reg.Nombre)})
}

// Función para ejecutar una función registrada
func ejecutarFuncion(w http.ResponseWriter, r *http.Request) {
	nombre := r.URL.Query().Get("nombre")

	mu.Lock()
	defer mu.Unlock()

	funcion, existe := funcionesRegistradas[nombre]
	if !existe {
		http.Error(w, "Función no encontrada", http.StatusNotFound)
		return
	}

	// Ejecutamos la función y obtenemos el resultado
	resultado := funcion()
	json.NewEncoder(w).Encode(map[string]interface{}{"resultado": resultado})
}

func main() {
	http.HandleFunc("/registrar", registrarFuncion)
	http.HandleFunc("/ejecutar", ejecutarFuncion)
	
	fmt.Println("Servidor escuchando en :5000 . Prueba 1")
	http.ListenAndServe(":5000", nil)
}
