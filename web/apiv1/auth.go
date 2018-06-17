package apiv1

import (
    "time"
    "io/ioutil"
    "encoding/json"
    "net/http"

    jwt "github.com/dgrijalva/jwt-go"

    "github.com/embik/pfennig/web/auth"
)

var SigningKey = []byte("AllYourBase")

type ApiClaims struct {
    Username string     `json:"user"`
    UserID  uint        `json:"user_id"`
    jwt.StandardClaims
}

type ReturnToken struct {
    Token   string      `json:"token"`
    Expiry  int64       `json:"expiry"`
}

type LoginMessage struct {
    Username    string  `json:"username"`
    Password    string  `json:"password"`
}

func GetToken(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    var msg LoginMessage
    err = json.Unmarshal(body, &msg)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    if isValid, user := auth.ValidateLogin(msg.Username, msg.Password); isValid {
        expiry := time.Now().AddDate(0, 0, 30)

        claims := ApiClaims{
            Username: user.Username,
            UserID: user.ID,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt:  expiry.Unix(),
                Issuer:     "pfennig/v1",
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        ss, _ := token.SignedString(SigningKey)

        json.NewEncoder(w).Encode(Response{
            Success: true,
            ErrMsg: "",
            ApiEndpoint: "v1/get_token",
            Payload: ReturnToken{Token: ss, Expiry: expiry.Unix()},
        })
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{
            Success: false,
            ErrMsg: "User or Passwort incorrect",
            ApiEndpoint: "v1/get_token",
            Payload: "",
        })
    }
}
