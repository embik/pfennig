package apiv1

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/embik/pfennig/app"
)

func GetAccountTypes(w http.ResponseWriter, r *http.Request) {
    user, _ := r.Context().Value("userID").(int)
    types := app.GetAccountTypes(user)
    json.NewEncoder(w).Encode(types)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("userID")
    fmt.Println(user)
    json.NewEncoder(w).Encode([]int{0})
}
