package models

type AccountType struct {
    ID          uint    `json:"id"`
    Label       string  `json:"label"`
    UserID      int     `json:"-"`
    IsGlobal    bool    `json:"is_global"`
}

type Account struct {
    Name        string      `json:"name"`
    Bank        string      `json:"bank"`
    AccountTypeID uint      `json:"account_type_id"`
}
