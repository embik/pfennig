package apiv1

import (
    "encoding/json"
    "net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(nil)
}
