package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/horlaarsco/bank/src/models"
	"github.com/horlaarsco/bank/src/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.IUser
	json.NewDecoder(r.Body).Decode(&user)

	validateBody := user.ValidateBody()

	if validateBody != nil {
		utils.ErrorHandler(w, validateBody, http.StatusBadRequest)
		return
	}

	password := utils.Hash(user.Password)
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Password = password
	err := models.CreateUser(&user)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}
	utils.SuccessHandler(w, user, "User Created", http.StatusCreated, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data models.LoginDTO
	json.NewDecoder(r.Body).Decode(&data)

	validateBody := data.ValidateBody()

	if validateBody != nil {
		utils.ErrorHandler(w, validateBody, http.StatusBadRequest)
		return
	}

	var user models.IUser
	err := models.GetUserByEmail(data.Email, &user)

	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	verifyPassword := utils.Compare(user.Password, data.Password)
	if !verifyPassword {
		utils.ErrorHandler(w, errors.New("Password do not match"), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	generateToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	user.Token = generateToken

	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}
	utils.SuccessHandler(w, user, "Login Successfull", http.StatusCreated, nil)
}
