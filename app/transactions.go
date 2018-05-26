package app

import (
    "github.com/jinzhu/gorm"
)

type TransactionType uint

const (
    Expense TransactionType = 0
    Income  TransactionType = 1
)

type Transaction struct {
    gorm.Model
    User        User
    UserID      uint
    Account     Account
    AccountID   uint
    Type        TransactionType `json:"type"`
    Amount      float32         `json:"amount"`
    Currency    Currency        `json:"currency"`
}
