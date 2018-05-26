package db

import (
    "github.com/embik/pfennig/app"
)

func GetCategories() []app.Category {
    return []app.Category{app.Category{ID: 1, Name: "Apartment", Account: 1}}
}
