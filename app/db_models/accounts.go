package db_models

import (
    "github.com/jinzhu/gorm"
)

type AccountType struct {
    gorm.Model
    Label       string
    UserID      int
    User        User
    IsGlobal    bool
}

type Account struct {
    gorm.Model
    Name            string
    Bank            string
    AccountType     AccountType
    AccountTypeID   uint
    Users           []*User     `gorm:"many2many:user_accounts;"`
}
