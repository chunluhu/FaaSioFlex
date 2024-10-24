package services

import (
    "auth-system/models"
    "errors"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var users = []models.User{} // Simulación de base de datos

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func SaveUser(user models.User) error {
    users = append(users, user) // Simulación de guardado
    return nil
}

func AuthenticateUser(email, password, phone string) (*models.User, error) {
    for _, user := range users {
        if user.Email == email && user.Phone == phone && CheckPasswordHash(password, user.Password) {
            return &user, nil
        }
    }
    return nil, errors.New("usuario no encontrado o credenciales incorrectas")
}

func GenerateJWT(user *models.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    user.ID,
        "email": user.Email,
        "exp":   time.Now().Add(time.Hour * 1).Unix(),
    })

    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
