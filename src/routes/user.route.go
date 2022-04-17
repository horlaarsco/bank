package routes

import (
	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/controllers"
)

func UserRoutes(mux *mux.Router) {
	mux.HandleFunc("/user", controllers.GetUser).Methods("GET")
	mux.HandleFunc("/user", controllers.CreateUser).Methods("POST")
}
