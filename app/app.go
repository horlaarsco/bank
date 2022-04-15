package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/customers", GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)
	router.HandleFunc("/customers", postCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getCustomer(res http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(res, vars["customer_id"])
}

func postCustomer(res http.ResponseWriter, request *http.Request) {
	fmt.Fprint(res, "Post request received")
}
