package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/guilhermecoelho/kakeibo/configurations"
)

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	Token string `json:"token"`
}

func Login(resp http.ResponseWriter, r *http.Request) {

	var authdetails Authentication
	var authResponse AuthenticationResponse

	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		resp.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(err)
		return
	}
	//user := GetUsers()
	token, _ := configurations.NewToken()
	authResponse.Token = token

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(authResponse)
}
