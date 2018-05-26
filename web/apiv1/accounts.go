package apiv1

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/embik/pfennig/app/db"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
    if user, ok := r.Context().Value("user").(int); ok {
        fmt.Println(user)
    }
    json.NewEncoder(w).Encode(db.Accounts)
}
