package middleware

import (
	"log"
	"net/http"
)

func NothingLocal(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[NOTHING] from local %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
