package app

import (
	"github.com/embik/pfennig/app/models"
    "github.com/embik/pfennig/app/db_models"
)

func GetAccountTypes(userID int) []models.AccountType {
    var account_types []db_models.AccountType

    db := getDB()
    db.Where(&db_models.AccountType{UserID: userID}).Or(&db_models.AccountType{IsGlobal: true}).Find(&account_types)
    return convertAccountTypes(account_types)
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
