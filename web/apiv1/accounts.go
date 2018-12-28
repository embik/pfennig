package apiv1

import (
    "io/ioutil"
    "encoding/json"
    "net/http"

    "github.com/embik/pfennig/pkg/data"
    "github.com/embik/pfennig/pkg/data/models"
)

func GetAccountTypes(w http.ResponseWriter, r *http.Request) {
    user := uint(r.Context().Value("userID").(float64))
    types := data.GetAccountTypes(user)

    json.NewEncoder(w).Encode(Response{
        Success: true,
        ErrMsg: "",
        ApiEndpoint: "v1/get_account_types",
        Payload: types,
    })
}

func CreateAccountType(w http.ResponseWriter, r *http.Request) {
    ok, user := data.GetUserByID(uint(r.Context().Value("userID").(float64)))
    if !ok || !user.IsAdmin {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{
            Success: false,
            ErrMsg: "You're not authorized to create new account types",
            ApiEndpoint: "v1/create_account_type",
            Payload: "",
        })
        return
    }

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
    }

    if data.CreateAccountType(account_type) {
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
    accounts := data.GetAccounts(user)
    json.NewEncoder(w).Encode(Response{
        Success: true,
        ErrMsg: "",
        ApiEndpoint: "v1/get_accounts",
        Payload: accounts,
    })
}
