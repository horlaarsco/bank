package utils

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Hash(passsword string) string {
	password := []byte("MyDarkSecret")
	pass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(pass)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
