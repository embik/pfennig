package data

import (
	"github.com/embik/pfennig/pkg/data/models"
    "github.com/embik/pfennig/pkg/data/db_models"
)

func CreateAccount(data models.Account) bool {
    var users []*db_models.User
    err := getDB().Create(&db_models.Account{
        Name:           data.Name,
        Bank:           data.Bank,
        AccountTypeID:  data.AccountTypeID,
        Users:          users,
    }).Error

    return err == nil
}

func GetAccounts(userID uint) []models.Account {
    var accounts []db_models.Account
    var user db_models.User
    getDB().First(&user, userID)
    getDB().Model(&user).Related(&accounts, "Accounts")
    getDB().Preload("Users").First(&accounts)

    return convertAccounts(accounts)
}

func GetAccountTypes(userID uint) []models.AccountType {
    var account_types []db_models.AccountType
    getDB().Find(&account_types)

    return convertAccountTypes(account_types)
}

func CreateAccountType(data models.AccountType) bool {
    err := getDB().Create(&db_models.AccountType{
        Label:      data.Label,
    }).Error

    return err == nil
}

func convertAccounts(data []db_models.Account) []models.Account {
    var accounts []models.Account
    for _, e := range data {
        var user_ids []uint
        for _, user := range e.Users {
            user_ids = append(user_ids, user.ID)
        }
        accounts = append(accounts, models.Account{
            ID: e.ID,
            Name: e.Name,
            Bank: e.Bank,
            AccountTypeID: e.AccountTypeID,
            UserIDs: user_ids,
        })
    }
    return accounts
}

func convertAccountTypes(data []db_models.AccountType) []models.AccountType {
    var account_types []models.AccountType
    for _, e := range data {
        account_types = append(account_types, models.AccountType{
            ID: e.ID,
            Label: e.Label,
        })
    }

    return account_types
}
