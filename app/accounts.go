package app

import (
	"github.com/embik/pfennig/app/models"
    "github.com/embik/pfennig/app/db_models"
)

func GetAccountTypes(userID uint) []models.AccountType {
    var account_types []db_models.AccountType

    db := getDB()
    db.Where(&db_models.AccountType{UserID: int(userID)}).Or(&db_models.AccountType{IsGlobal: true}).Find(&account_types)
    return convertAccountTypes(account_types)
}

func CreateAccountType(data models.AccountType) bool {
    err := getDB().Create(&db_models.AccountType{
        Label:      data.Label,
        IsGlobal:   data.IsGlobal,
        UserID:     data.UserID,
    }).Error

    return err == nil
}

func convertAccountTypes(data []db_models.AccountType) []models.AccountType {
    var account_types []models.AccountType
    for _, e := range data {
        account_types = append(account_types, models.AccountType{
            ID: e.ID,
            Label: e.Label,
            UserID: e.UserID,
            IsGlobal: e.IsGlobal,
        })
    }

    return account_types
}
