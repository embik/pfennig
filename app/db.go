package app

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/embik/pfennig/app/db_models"
)

var database *gorm.DB

func InitDB(databasePath string) error {
    var err error
    database, err = gorm.Open("sqlite3", databasePath)
    if err == nil {
        database.AutoMigrate(&db_models.User{})
        database.AutoMigrate(&db_models.AccountType{})
        database.AutoMigrate(&db_models.Account{})

        createDefaultData()
    }
    return err
}

func createDefaultData() {
    database.FirstOrCreate(&db_models.AccountType{Label: "Girokonto"})
}

func CloseDB() {
    database.Close()
}

func getDB() *gorm.DB {
    return database
}
