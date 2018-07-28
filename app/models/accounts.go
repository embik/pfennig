package models

type AccountType struct {
    ID          uint    `json:"id"`
    Label       string  `json:"label"`
}

type Account struct {
    ID          uint        `json:"id"`
    Name        string      `json:"name"`
    Bank        string      `json:"bank"`
    AccountTypeID uint      `json:"account_type_id"`
    UserIDs     []uint      `json:"user_ids"`
}
