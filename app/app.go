package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horlaarsco/bank/domain"
	"github.com/horlaarsco/bank/service"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}
