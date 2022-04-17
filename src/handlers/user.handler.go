package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/horlaarsco/bank/src/models"
	"github.com/horlaarsco/bank/src/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id uint = 1
	var user models.IUser
	err := models.GetUser(id, &user)
	if err != "" {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.IUser
	json.NewDecoder(r.Body).Decode(&user)

	password := utils.Hash(user.Password)
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Password = password
	err := models.CreateUser(user)
	if err != "" {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		response, _ := json.Marshal(err)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
