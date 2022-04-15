package app

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	name string
	age  int
}

func GetAllCustomers(res http.ResponseWriter, request *http.Request) {
	customers := []Customer{{"John", 30}, {"Jane", 25}}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(customers)
}
