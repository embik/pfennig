package db_models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username    string      `gorm:"unique"`
    FirstName   string
    LastName    string
    Email       string
    PwdHash     string
    IsAdmin     bool
    Accounts    []*Account  `gorm:"many2many:user_accounts;"`
}
