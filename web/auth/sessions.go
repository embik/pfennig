package auth

import (
    "golang.org/x/crypto/bcrypt"

    "github.com/embik/pfennig/pkg/data"
    "github.com/embik/pfennig/pkg/data/models"
)

func ValidateLogin(username string, password string) (bool, models.User) {
    var ok bool
    var user models.User

    if ok, user = data.GetUser(username); ok {
        if bcrypt.CompareHashAndPassword([]byte(user.PwdHash), []byte(password)) == nil {
            return true, user
        }
    }

    return false, user
}
