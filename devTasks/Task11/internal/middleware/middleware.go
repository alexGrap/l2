package middleware

import (
	"log"
	"net/http"
	"time"
)

func MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println(
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
		)
		next(w, r)
	}
}
