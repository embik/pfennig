package apiv1

import (
    "encoding/json"
    "net/http"

    "github.com/embik/pfennig/app/db"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
    categories := db.GetCategories()
    json.NewEncoder(w).Encode(categories)
}
