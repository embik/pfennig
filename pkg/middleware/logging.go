package middleware

import (
    "net/http"

    log "github.com/sirupsen/logrus"
)

func RequestLogging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
        log.WithFields(log.Fields{
            "proto": r.Proto,
            "method": r.Method,
            "host": r.Host,
            "url": r.URL,
            "remote": r.RemoteAddr,
        }).Info("received http request")
    })
}
