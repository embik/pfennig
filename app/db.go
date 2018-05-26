package app

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB(dbPath string) error {
    var err error
    db, err = gorm.Open("sqlite3", dbPath)
    if err == nil {
        db.AutoMigrate(&User{})
        db.AutoMigrate(&AccountType{})
        db.AutoMigrate(&Account{})

        createDefaultData()
    }
    return err
}

func createDefaultData() {
    db.FirstOrCreate(&AccountType{Label: "Girokonto"})
}

func CloseDB() {
    db.Close()
}

func GetDB() *gorm.DB {
    return db
}
