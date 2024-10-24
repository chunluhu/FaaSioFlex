package services

import (
    "errors"
    "your_project_name/models"
)

var users []models.User

func Register(user models.User) (string, error) {
    // Aquí debería agregar lógica para verificar si el usuario ya existe
    // Por ejemplo, verificar si el email o el nombre de usuario ya están en uso y el telefono porque 
    //cambiar el email es muy sencillo
    for _, existingUser := range users {
        if existingUser.Username == user.Username || existingUser.Email == user.Email || existingUser.Phone == user.Phone {
            return "", errors.New("El nombre de usuario o el email  o Phone ya están en uso")
        }
    }
    // Aquí deberías agregar lógica para verificar si el usuario ya existe
    users = append(users, user)
    token := generateToken(user) // Genera un token de autenticación
    return token, nil
}

func Authenticate(username, password string) (string, error) {
    for _, user := range users {
        if (user.Username == identifier || user.Email == identifier) && user.Password == password {
            token := generateToken(user) // Genera un token de autenticación
            return token, nil
        }
    }
    return "", errors.New("Credenciales incorrectas")
}

func generateToken(user models.User) string {
    // Implementa la lógica para generar un token (puedes usar JWT)
    return "token_placeholder" // Reemplaza con un token real
}
