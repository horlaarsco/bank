package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Hash(passsword string) string {
	password := []byte("MyDarkSecret")
	pass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(pass)
}
