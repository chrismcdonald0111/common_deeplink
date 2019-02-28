package service

import (
    "net/http"
)

func (s *server) redirect() http.HandlerFunc {
    // do something
    return func(w http.ResponseWriter, r *http.Request) {
         // See HTTP status codes here: https://golang.org/src/net/http/status.go
         http.Redirect(w, r, "http://www.donothingfor2minutes.com", http.StatusSeeOther)
    }
}
