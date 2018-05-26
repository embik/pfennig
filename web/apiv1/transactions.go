package apiv1

import (
    "encoding/json"
    "net/http"

    "github.com/embik/pfennig/app/db"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(db.Transactions)
}
