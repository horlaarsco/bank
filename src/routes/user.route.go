package routes

import (
	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/handlers"
)

func UserRoutes(mux *mux.Router) {
	mux.HandleFunc("/users", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/users", handlers.CreateUser).Methods("POST")
}
