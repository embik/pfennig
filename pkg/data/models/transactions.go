package models

type TransactionType uint

const (
    Expense TransactionType = 0
    Income  TransactionType = 1
)

type Transaction struct {
    User        User
    UserID      uint
    Account     Account
    AccountID   uint
    Type        TransactionType `json:"type"`
    Amount      float32         `json:"amount"`
    Currency    Currency        `json:"currency"`
}
