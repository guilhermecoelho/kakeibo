package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilhermecoelho/kakeibo/data"
	"github.com/guilhermecoelho/kakeibo/models"
)

var income = models.Income{}

func GetIncomes(resp http.ResponseWriter, r *http.Request) {

	incomes, errorData := data.GetIncomes()

	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := incomes.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func GetIncomeById(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}

	income, errorData := data.GetIncomesById(id)
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := income.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func PutIncome(resp http.ResponseWriter, r *http.Request) {

	errDecode := income.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PutIncomes(&income)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func PostIncome(resp http.ResponseWriter, r *http.Request) {

	errDecode := income.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PostIncomes(&income)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func DeleteIncome(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}
	income, errGet := data.GetIncomesById(id)
	if errGet != nil {
		http.Error(resp, errGet.Error(), http.StatusBadRequest)
		return
	}

	errorProcess := data.DeleteIncomes(&income)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
	}
}
