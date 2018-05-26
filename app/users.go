package app

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username    string  `json:"username",gorm:"unique_index"`
    FirstName   string  `json:"first_name"`
    LastName    string  `json:"last_name"`
    Email       string  `json:"email"`
    PwdHash     string  `json:"-"`
}
