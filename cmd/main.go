package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/routes"
)

func main() {
	app := mux.NewRouter()
	routes.RegisterBookRoutes(app)
	http.Handle("/", app)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
