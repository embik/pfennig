package handlers

import (
    "net/http"

    "github.com/gorilla/csrf"

    "github.com/embik/pfennig/web/tmpl"
    "github.com/embik/pfennig/web/auth"
)

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
    data := make(map[string]interface{})
    data["Session"] = auth.GetSession(r)
    data[csrf.TemplateTag] = csrf.TemplateField(r)

    tmpl.Templates["index.html"].ExecuteTemplate(w, "base", data)
}
