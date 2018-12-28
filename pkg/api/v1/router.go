package apiv1

import (
    "net/http"
    "github.com/gorilla/mux"
)

func WireUpRoutes(router *mux.Router) {
    r := (*router).PathPrefix("/v1/").Subrouter()

    r.HandleFunc("/login", GetToken).Methods("POST")
    r.Handle("/account_types", requireToken(http.HandlerFunc(GetAccountTypes))).Methods("GET")
    r.Handle("/account_types", requireToken(http.HandlerFunc(CreateAccountType))).Methods("PUT")
    r.Handle("/accounts", requireToken(http.HandlerFunc(GetAccounts))).Methods("GET")
    r.HandleFunc("/users", GetUsers).Methods("GET")
    r.HandleFunc("/transactions", GetTransactions).Methods("GET")
    r.HandleFunc("/categories", GetCategories).Methods("GET")

    r.Use(asJSONMiddleware)
}
