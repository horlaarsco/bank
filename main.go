package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/src/routes"
	"github.com/horlaarsco/bank/src/utils"
)

func main() {
	PORT := 8080
	app := mux.NewRouter()
	app.Use(utils.LoggingMiddleware)

	routes.Auth(app)
	routes.User(app)

	http.Handle("/", app)

	fmt.Printf("Server is running on port %v\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), app))
}
