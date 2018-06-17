package auth

import (
    "net/http"

    "golang.org/x/crypto/bcrypt"
    "github.com/gorilla/sessions"

    "github.com/embik/pfennig/app"
    "github.com/embik/pfennig/app/models"
)

var (
    key = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)
)

type Session struct {
    IsAuthenticated bool
    Username        string
}

func GetSession(r *http.Request) Session {
    session, err := store.Get(r, "pfennig-session")
    if err != nil {
        return Session{IsAuthenticated: false, Username: ""}
    }

    isAuthenticated := false
    if session.Values["authenticated"] != nil {
        isAuthenticated = session.Values["authenticated"].(bool)
    }

    username := ""
    if session.Values["username"] != nil {
        username = session.Values["username"].(string)
    }

    return Session{
        IsAuthenticated: isAuthenticated,
        Username: username,
    }
}

func SignIn(w http.ResponseWriter, r *http.Request, username string, password string) {
    session, _ := store.Get(r, "pfennig-session")

    if isValid, user := ValidateLogin(username, password); isValid {
        session.Values["authenticated"] = true
        session.Values["username"] = user.Username
    } else {
        session.Values["authenticated"] = false
        session.Values["username"] = ""
    }

    session.Save(r, w)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "pfennig-session")

    session.Values["authenticated"] = false
    session.Values["username"] = ""

    session.Save(r, w)
}

func ValidateLogin(username string, password string) (bool, models.User) {
    var ok bool
    var user models.User

    if ok, user = app.GetUser(username); ok {
        if bcrypt.CompareHashAndPassword([]byte(user.PwdHash), []byte(password)) == nil {
            return true, user
        }
    }

    return false, user
}
