package tmpl

import (
    "html/template"
    "github.com/embik/pfennig/web/router"
)

var Templates map[string]*template.Template

func InitTemplates() {
    tmplMap := template.FuncMap{
        "getURI": router.GetURI,
    }

    Templates = make(map[string]*template.Template)
    Templates["index.html"] = template.Must(template.New("").Funcs(tmplMap).ParseFiles("templates/base.html", "templates/index.html"))
}
