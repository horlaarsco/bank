package routes

import (
	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/handlers"
)

func Auth(mux *mux.Router) {
	mux.HandleFunc("/auth/signup", handlers.SignUp).Methods("POST")
	mux.HandleFunc("/auth/login", handlers.Login).Methods("POST")
}
