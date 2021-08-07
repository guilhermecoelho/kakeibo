package handlers

import (
	"net/http"

	"github.com/guilhermecoelho/kakeibo/data"
	"github.com/guilhermecoelho/kakeibo/models"
)

var reportReq = models.ReportRequest{}
var report = models.Report{}

type testIncome struct {
}

func (testIncome) Find() (models.Incomes, error) {
	incomes := models.Incomes{}

	income := models.Income{}
	income.Id = 1
	income.Amount = 10

	income2 := models.Income{}
	income2.Id = 2
	income2.Amount = 15

	incomes = append(incomes, &income)
	incomes = append(incomes, &income2)

	return incomes, nil
}

func GetReport(resp http.ResponseWriter, r *http.Request) {

	// report, errorData := data.GetCategories()
	// if errorData != nil {
	// 	http.Error(resp, errorData.Error(), http.StatusInternalServerError)
	// 	return
	// }
	errDecode := reportReq.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	incomes, _ := data.GetIncomesTest(testIncome{})

	report.DateStart = reportReq.DateStart
	report.DateFinish = reportReq.DateFinish

	expense := models.Expense{}
	expense.Id = 1
	expense.Amount = 40
	report.Expenses = append(report.Expenses, &expense)

	// income := models.Income{}
	// income.Id = 1
	// income.Amount = 10

	// income2 := models.Income{}
	// income2.Id = 2
	// income2.Amount = 15

	report.Incomes = incomes

	for i := 0; i < len(report.Incomes); i++ {
		report.TotalIncome = report.TotalIncome + report.Incomes[i].Amount
	}

	for i := 0; i < len(report.Expenses); i++ {
		report.TotalExpense = report.TotalExpense + report.Expenses[i].Amount
	}

	report.Balance = report.TotalIncome - report.TotalExpense

	err := report.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}
