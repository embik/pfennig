package app

import (
    "github.com/embik/pfennig/app/models"
    "github.com/embik/pfennig/app/db_models"
)

func GetUser(username string) (bool, models.User) {
    var db_user db_models.User
    var user    models.User

    if getDB().Where(&db_models.User{Username: username}).First(&db_user).RecordNotFound() {
        return false, user
    } else {
        user = convertUser(db_user)
        return true, user
    }
}

func convertUser(data db_models.User) models.User {
    return models.User{
        ID: data.ID,
        Username: data.Username,
        FirstName: data.FirstName,
        LastName: data.LastName,
        Email: data.Email,
        PwdHash: data.PwdHash,
    }
}

