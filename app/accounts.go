package app

import (
    "github.com/jinzhu/gorm"
)

type AccountType struct {
    gorm.Model
    Label   string      `gorm:"unique"`
}

type Account struct {
    gorm.Model
    Name        string      `json:"name"`
    Bank        string      `json:"bank"`
    AccountType AccountType `json:"-"`
    AccountTypeID uint      `json:"account_type_id"`
}
