package apiv1

import (
    "encoding/json"
    "net/http"

    "github.com/embik/pfennig/pkg/data/models"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
    var categories []models.Category
    json.NewEncoder(w).Encode(categories)
}
