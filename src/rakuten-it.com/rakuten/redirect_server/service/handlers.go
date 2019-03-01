package service

import (
    "net/http"
)

// Redirect to new URL
func (s *server) redirect() http.HandlerFunc {
    // Execute things above closure exactly once, before server start-up
    return func(w http.ResponseWriter, r *http.Request) {
         // See HTTP status codes here: https://golang.org/src/net/http/status.go
         http.Redirect(w, r, "http://www.donothingfor2minutes.com", http.StatusSeeOther)
    }
}
