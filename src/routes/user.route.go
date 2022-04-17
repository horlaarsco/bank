package routes

import (
	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/handlers"
)

func UserRoutes(mux *mux.Router) {
	mux.HandleFunc("/user", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/user", handlers.CreateUser).Methods("POST")
}
