package apiv1

import (
    "io/ioutil"
    "encoding/json"
    "net/http"

    "github.com/embik/pfennig/app"
    "github.com/embik/pfennig/app/models"
)

func GetAccountTypes(w http.ResponseWriter, r *http.Request) {
    user := uint(r.Context().Value("userID").(float64))
    types := app.GetAccountTypes(user)

    json.NewEncoder(w).Encode(Response{
        Success: true,
        ErrMsg: "",
        ApiEndpoint: "v1/get_account_types",
        Payload: types,
    })
}

func CreateAccountType(w http.ResponseWriter, r *http.Request) {
    user := uint(r.Context().Value("userID").(float64))

    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    var request AccountTypeRequest
    if json.Unmarshal(body, &request) != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(Response{
            Success: false,
            ErrMsg: "Invalid Data",
            ApiEndpoint: "v1/create_account_type",
            Payload: "",
        })
        return
    }

    account_type := models.AccountType{
        Label: request.Label,
        IsGlobal: request.IsGlobal,
        UserID: int(user),
    }

    if app.CreateAccountType(account_type) {
        json.NewEncoder(w).Encode(Response{
            Success: true,
            ErrMsg: "",
            ApiEndpoint: "v1/create_account_type",
            Payload: "",
        })
    } else {
        json.NewEncoder(w).Encode(Response{
            Success: false,
            ErrMsg: "Creating Account Type failed",
            ApiEndpoint: "v1/create_account_type",
            Payload: "",
        })
    }
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
    user := uint(r.Context().Value("userID").(float64))
    accounts := app.GetAccounts(user)
    json.NewEncoder(w).Encode(Response{
        Success: true,
        ErrMsg: "",
        ApiEndpoint: "v1/get_accounts",
        Payload: accounts,
    })
}
