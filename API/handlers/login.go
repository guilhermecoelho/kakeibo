package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/data"
	"github.com/guilhermecoelho/kakeibo/models"
)

type AuthenticationResponse struct {
	Token string `json:"token"`
}

var login = models.Login{}

func Login(resp http.ResponseWriter, r *http.Request) {

	var authResponse AuthenticationResponse

	errDecode := login.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	user, err := data.GetUserByName(login.User)
	if err != nil {
		http.Error(resp, "invalid password or user", http.StatusBadRequest)
		return
	}

	compareHash, errorHash := configurations.CompareHashAndPassword(login.Password, user.Hash)
	if errorHash != nil {
		http.Error(resp, "invalid password or user", http.StatusBadRequest)
		return
	}

	if !compareHash {
		http.Error(resp, "invalid password or user", http.StatusBadRequest)
		return
	}

	token, _ := configurations.NewToken()
	authResponse.Token = token

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(authResponse)
}
