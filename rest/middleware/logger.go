package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s%s from %s", r.Method, r.URL.Path, r.URL.RawQuery, r.RemoteAddr)

		next.ServeHTTP(w, r)
		
		log.Printf("Completed %s%s in %v", r.URL.Path, r.URL.RawQuery, time.Since(start))
	}

	nextHandler := http.HandlerFunc(handler)
	return nextHandler
}
