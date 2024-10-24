package models

type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Phone    string `json:"phone"`
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}
