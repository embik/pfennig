package router

import (
    "log"
    "github.com/gorilla/mux"
)

var (
    r = mux.NewRouter()
)

func GetRouter() *mux.Router {
    return r
}

func GetURI(name string, params ...string) string {
    url, err := r.Get(name).URL(params...)
    if err != nil {
        log.Fatal("Failed Building Route")
    }
    return url.RequestURI()
}

