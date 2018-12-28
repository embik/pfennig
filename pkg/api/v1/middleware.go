package apiv1

import (
	"fmt"
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func requireToken(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("X-API-Token")
        if tokenString != "" {
            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
                }

                return SigningKey, nil
            })

            if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
                context := context.WithValue(r.Context(), "userID", claims["user_id"])
                next.ServeHTTP(w, r.WithContext(context))
            } else {
                fmt.Printf("Error: %s\n", err)
            }
        }
    })
}
