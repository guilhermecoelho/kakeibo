package services

import (
	"github.com/guilhermecoelho/kakeibo/data"
	"github.com/guilhermecoelho/kakeibo/models"
)

type ReportInterface interface {
	FindIncome() (models.Incomes, error)
	FindExpense() (models.Expenses, error)
	GetDates() [2]string
}

type ReportRequest struct {
	DateStart  string
	DateFinish string
}

func (r ReportRequest) FindIncome() (models.Incomes, error) {

	incomes, errorData := data.GetIncomesByPeriod(r.DateStart, r.DateFinish)
	if errorData != nil {
		return incomes, errorData
	}
	return incomes, nil
}

func (r ReportRequest) FindExpense() (models.Expenses, error) {

	expenses, errorData := data.GetExpensesByPeriod(r.DateStart, r.DateFinish)
	if errorData != nil {
		return expenses, errorData
	}
	return expenses, nil
}

func (r ReportRequest) GetDates() [2]string {

	date := [...]string{r.DateStart, r.DateFinish}

	return date
}

func GetReport(r ReportInterface) (models.Report, error) {

	var report = models.Report{}

	dates := r.GetDates()

	report.DateStart = dates[0]
	report.DateFinish = dates[1]

	respIncome, _ := r.FindIncome()
	report.Incomes = respIncome

	respExpense, _ := r.FindExpense()
	report.Expenses = respExpense

	for i := 0; i < len(report.Incomes); i++ {
		report.TotalIncome = report.TotalIncome + report.Incomes[i].Amount
	}

	for i := 0; i < len(report.Expenses); i++ {
		report.TotalExpense = report.TotalExpense + report.Expenses[i].Amount
	}

	report.Balance = report.TotalIncome - report.TotalExpense

	return report, nil
}
