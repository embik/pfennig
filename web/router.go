package web

import (
    "github.com/gorilla/mux"

    "github.com/embik/pfennig/web/router"
    v1 "github.com/embik/pfennig/web/apiv1"
)

func NewRouter() *mux.Router {
    r := router.GetRouter()

    apiRouter := r.PathPrefix("/api/").Subrouter()
    v1.WireUpRoutes(apiRouter)

    return r
}
