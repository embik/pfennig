package app

import (
    "golang.org/x/crypto/bcrypt"
)

func CreateDummyData() {
    bytes, _ := bcrypt.GenerateFromPassword([]byte("testtest"), 14)
    db.FirstOrCreate(&User{Username: "embik", FirstName: "Marvin", LastName: "Beckers", Email: "mail@embik.me", PwdHash: string(bytes)})
}
