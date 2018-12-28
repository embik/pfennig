package apiv1

import (
    "encoding/json"
    "net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(nil)
}
