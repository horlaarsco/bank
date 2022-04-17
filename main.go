package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/routes"
)

func main() {
	PORT := 8080
	app := mux.NewRouter()
	routes.UserRoutes(app)
	http.Handle("/", app)

	fmt.Printf("Server is running on port %v\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), app))
}
