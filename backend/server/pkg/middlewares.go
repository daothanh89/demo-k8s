package pkg

import (
	"log"
	"net/http"
	"strings"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] method=%s uri=%s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorization!"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		if ok := verifyToken(tokenString); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying token"))
			return
		}

		next.ServeHTTP(w, r)
	}
}
