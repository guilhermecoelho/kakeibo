package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilhermecoelho/kakeibo/configurations"
	"github.com/guilhermecoelho/kakeibo/data"
	"github.com/guilhermecoelho/kakeibo/models"
)

var user = models.User{}

func GetUsers(resp http.ResponseWriter, r *http.Request) {

	users, errorData := data.GetUsers()
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := users.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func GetUserById(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}

	user, errorData := data.GetUserById(id)
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := user.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func GetUserByName(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	name := params["name"]
	if name == "" {
		http.Error(resp, "invalid user name", http.StatusBadRequest)
		return
	}

	user, errorData := data.GetUserByName(name)
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := user.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func PutUser(resp http.ResponseWriter, r *http.Request) {

	errDecode := user.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PutUser(&user)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func PostUser(resp http.ResponseWriter, r *http.Request) {

	errDecode := user.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	generateHash, errGenerate := configurations.GenerateHash(user.Password)
	if errGenerate != nil {
		http.Error(resp, errGenerate.Error(), http.StatusInternalServerError)
		return
	}
	user.Hash = generateHash

	dataUser, errorProcess := data.PostUser(&user)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataUser.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func DeleteUser(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}
	user, errGet := data.GetUserById(id)
	if errGet != nil {
		http.Error(resp, errGet.Error(), http.StatusBadRequest)
		return
	}

	errorProcess := data.DeleteUser(&user)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
	}
}
