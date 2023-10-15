package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s Method: %s", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
