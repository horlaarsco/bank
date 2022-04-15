package app

import (
	"encoding/json"
	"net/http"

	"github.com/horlaarsco/bank/service"
)

type Customer struct {
	name string
	age  int
}

type CustomerHandlers struct {
	service service.CustomerServive
}

func (ch *CustomerHandlers) getAllCustomers(res http.ResponseWriter, request *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(customers)
}
