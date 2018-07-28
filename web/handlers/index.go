package handlers

import (
    "net/http"

    "github.com/embik/pfennig/web/tmpl"
)

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl.Templates["index.html"].ExecuteTemplate(w, "base", nil)
}
