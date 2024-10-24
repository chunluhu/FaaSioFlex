package controllers

import (
    "auth-system/models"
    "auth-system/services"
    "encoding/json"
    "net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    if user.Name == "" || user.Email == "" || user.Password == "" || user.Phone == "" {
        http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
        return
    }

    hashedPassword, err := services.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Error al generar hash", http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword

    err = services.SaveUser(user)
    if err != nil {
        http.Error(w, "Error al guardar usuario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Usuario registrado exitosamente")
}

func Login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
        Phone    string `json:"phone"`
    }

    _ = json.NewDecoder(r.Body).Decode(&credentials)

    user, err := services.AuthenticateUser(credentials.Email, credentials.Password, credentials.Phone)
    if err != nil {
        http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
        return
    }

    token, err := services.GenerateJWT(user)
    if err != nil {
        http.Error(w, "Error al generar token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}
