package models

type AccountType struct {
    ID          uint    `json:"id"`
    Label       string  `json:"label"`
    IsGlobal    bool    `json:"is_global"`
    UserID      int     `json:"-"`
}

type Account struct {
    Name        string      `json:"name"`
    Bank        string      `json:"bank"`
    AccountTypeID uint      `json:"account_type_id"`
}
