package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilhermecoelho/kakeibo/data"
	"github.com/guilhermecoelho/kakeibo/models"
)

var expense = models.Expense{}

func GetExpenses(resp http.ResponseWriter, r *http.Request) {

	expenses, errorData := data.GetExpenses()
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := expenses.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func GetExpenseById(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}

	expense, errorData := data.GetExpensesById(id)
	if errorData != nil {
		http.Error(resp, errorData.Error(), http.StatusInternalServerError)
		return
	}

	err := expense.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func PutExpense(resp http.ResponseWriter, r *http.Request) {

	errDecode := expense.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PutExpenses(&expense)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func PostExpense(resp http.ResponseWriter, r *http.Request) {

	errDecode := expense.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	dataMoviment, errorProcess := data.PostExpenses(&expense)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
		return
	}
	err := dataMoviment.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func DeleteExpense(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, errorReq := strconv.Atoi(params["id"])
	if errorReq != nil {
		http.Error(resp, errorReq.Error(), http.StatusBadRequest)
		return
	}
	expense, errGet := data.GetExpensesById(id)
	if errGet != nil {
		http.Error(resp, errGet.Error(), http.StatusBadRequest)
		return
	}

	errorProcess := data.DeleteExpenses(&expense)
	if errorProcess != nil {
		http.Error(resp, errorProcess.Error(), http.StatusInternalServerError)
	}
}
