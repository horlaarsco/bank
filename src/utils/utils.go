package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func Hash(passsword string) string {
	password := []byte("MyDarkSecret")
	pass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(pass)
}
