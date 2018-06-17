package app

import (
    "golang.org/x/crypto/bcrypt"
    
    "github.com/embik/pfennig/app/db_models"
)

func CreateDummyData() {
    bytes, _ := bcrypt.GenerateFromPassword([]byte("testtest"), 14)
    getDB().FirstOrCreate(&db_models.User{Username: "embik", FirstName: "Marvin", LastName: "Beckers", Email: "mail@embik.me", PwdHash: string(bytes)})
}
