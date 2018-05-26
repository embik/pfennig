package handlers

import (
    "net/http"

    "github.com/gorilla/csrf"

    "github.com/embik/pfennig/web/tmpl"
    "github.com/embik/pfennig/web/router"
    "github.com/embik/pfennig/web/auth"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
    session := auth.GetSession(r)
    if session.IsAuthenticated {
        http.Redirect(w, r, router.GetURI("index"), 302)
    }
    tmpl.Templates["login.html"].ExecuteTemplate(w, "base", map[string]interface{}{
        csrf.TemplateTag: csrf.TemplateField(r),
    })
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.Form.Get("username")
    password := r.Form.Get("password")

    auth.SignIn(w, r, username, password)

    if true {
        http.Redirect(w, r, router.GetURI("index"), 302)
    }
}

func GetLogout(w http.ResponseWriter, r *http.Request) {
    auth.SignOut(w, r)

    http.Redirect(w, r, router.GetURI("login"), 302)
}
