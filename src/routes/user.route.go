package routes

import (
	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/handlers"
	"github.com/horlaarsco/bank/src/utils"
)

func User(mux *mux.Router) {
	userRoutes := mux.PathPrefix("/user").Subrouter()
	userRoutes.Use(utils.AuthMiddleware)

	userRoutes.HandleFunc("/", handlers.GetUser).Methods("GET")
}
