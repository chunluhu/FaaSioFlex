 package routes

import (
    "auth-system/controllers"
    "github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
    r.HandleFunc("/register", controllers.Register).Methods("POST")
    r.HandleFunc("/login", controllers.Login).Methods("POST")
}

