package app

import (
    "golang.org/x/crypto/bcrypt"

    "github.com/embik/pfennig/app/db_models"
)

func CreateDummyData() {
    pwdEmbik, _ := bcrypt.GenerateFromPassword([]byte("testtest"), 14)
    var embik db_models.User

    getDB().Where(db_models.User{Username: "embik"}).Assign(db_models.User{
        FirstName: "Marvin",
        LastName: "Beckers",
        Email: "mail@embik.me",
        PwdHash: string(pwdEmbik),
    }).FirstOrCreate(&embik)

    pwdEmbik2, _ := bcrypt.GenerateFromPassword([]byte("testtest2"), 14)
    var embik2 db_models.User

    getDB().Where(db_models.User{Username: "embik2"}).Assign(db_models.User{
        FirstName: "Marvin",
        LastName: "Beckers",
        Email: "marvin@embik.me",
        PwdHash: string(pwdEmbik2),
    }).FirstOrCreate(&embik2)

    var account_1 db_models.Account
    getDB().Where(db_models.Account{Name: "Sparkonto"}).Assign(db_models.Account{
        Name:   "Sparkonto",
        Bank:   "Sparkasse",
        AccountTypeID:  1,
        Users:  []*db_models.User{&embik, &embik2},
    }).FirstOrCreate(&account_1)
}
