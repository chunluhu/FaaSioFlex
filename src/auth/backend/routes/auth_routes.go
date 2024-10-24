package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "your_project_name/handlers"
)

func SetupAuthRoutes(router *mux.Router) {
    router.HandleFunc("/api/register", handlers.RegisterUser).Methods("POST")
    router.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")
}

