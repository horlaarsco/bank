package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

var mySigningKey = []byte("secret")

func ErrorHandler(w http.ResponseWriter, err error, code int) {
	log.Println(err)
	response := Response{"error", err.Error(), nil, nil}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func SuccessHandler(w http.ResponseWriter, data interface{}, msg string, code int, meta interface{}) {
	response := Response{"success", msg, data, meta}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func Hash(passw string) string {
	password := []byte(passw)
	pass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(pass)
}

func Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ErrorHandler(w, errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		if strings.ToLower(strings.Split(authHeader, " ")[0]) != "bearer" {
			ErrorHandler(w, errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")[1]

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(bearerToken, claims,
			func(t *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ErrorHandler(w, errors.New("invalid token"), http.StatusUnauthorized)
				return
			}
			ErrorHandler(w, errors.New("bad request"), http.StatusBadRequest)
			return
		}

		if !token.Valid {
			ErrorHandler(w, errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		r.Header.Set("user_id", fmt.Sprintf("%d", claims.Id))
		next.ServeHTTP(w, r)
	})
}

func GenerateToken(id uint) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
