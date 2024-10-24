package models

import (
    "time"
)

type User struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Phone     string    `json:"phone"`  // Teléfono móvil
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
}
