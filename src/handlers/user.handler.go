package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/horlaarsco/bank/src/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var id string = r.Header.Get("user_id")

	userId, _ := strconv.ParseUint(id, 10, 32)

	var user models.IUser
	err := models.GetUser(uint(userId), &user)
	if err != "" {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
