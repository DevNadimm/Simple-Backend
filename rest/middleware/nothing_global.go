package middleware

import (
	"log"
	"net/http"
)

func NothingGlobal(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[NOTHING] from global %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
