package router

import (
    "log"
    "github.com/gorilla/mux"

    v1 "github.com/embik/pfennig/pkg/api/v1"
)

var (
    r = mux.NewRouter()
)

func GetRouter() *mux.Router {
    apiRouter := r.PathPrefix("/api/").Subrouter()
    v1.WireUpRoutes(apiRouter)
    return r
}

func GetURI(name string, params ...string) string {
    url, err := r.Get(name).URL(params...)
    if err != nil {
        log.Fatal("Failed Building Route")
    }
    return url.RequestURI()
}

