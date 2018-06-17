package models

type User struct {
    Username    string  `json:"username"`
    FirstName   string  `json:"first_name"`
    LastName    string  `json:"last_name"`
    Email       string  `json:"email"`
    PwdHash     string  `json:"-"`
}
