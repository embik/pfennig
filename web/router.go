package web

import (
    "net/http"

    "github.com/gorilla/mux"

    "github.com/embik/pfennig/web/router"
    "github.com/embik/pfennig/web/handlers"
    "github.com/embik/pfennig/web/auth"
    api "github.com/embik/pfennig/web/apiv1"
)

func NewRouter(assetPath string) *mux.Router {
    r := router.GetRouter()

    r.HandleFunc("/", requireLogin(handlers.GetIndexHandler)).Methods("GET").Name("index")

    r.HandleFunc("/login", handlers.GetLogin).Methods("GET").Name("login")
    r.HandleFunc("/login", handlers.PostLogin).Methods("POST")
    r.HandleFunc("/logout", requireLogin(handlers.GetLogout)).Methods("POST").Name("logout")


    staticFileDirectory := http.Dir(assetPath)
    staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
    r.PathPrefix("/static/").Handler(staticFileHandler)

    apiRouter := r.PathPrefix("/api/").Subrouter()
    api.WireUpRoutes(apiRouter)

    return r
}

func requireLogin(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        session := auth.GetSession(r)
        if session.IsAuthenticated == false {
            http.Redirect(w, r, router.GetURI("login"), 302)
        }

        next.ServeHTTP(w, r)
    }
}
