package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilhermecoelho/kakeibo/data"
	"github.com/guilhermecoelho/kakeibo/models"
)

var group = models.Group{}

func GetGroups(resp http.ResponseWriter, r *http.Request) {

	groups, errorData := data.GetGroups()
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := groups.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func GetGroupById(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}

	group, errorData := data.GetGroupById(id)
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := group.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func PutGroup(resp http.ResponseWriter, r *http.Request) {

	errDecode := group.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PutGroup(&group)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func PostGroup(resp http.ResponseWriter, r *http.Request) {

	errDecode := group.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PostGroup(&group)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func DeleteGroup(resp http.ResponseWriter, r *http.Request) {

	errDecode := group.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	errorProcess := data.DeleteGroup(&group)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
	}
}
