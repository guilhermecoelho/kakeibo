package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/guilhermecoelho/kakeibo/models"
	"github.com/guilhermecoelho/kakeibo/services"
)

type FakeReportRequest struct {
	DateStart  string
	DateFinish string
}

var incomes = models.Incomes{}
var expenses = models.Expenses{}

func populateIncome() {
	income := models.Income{}
	income.Id = 1
	income.Amount = 10

	income2 := models.Income{}
	income2.Id = 2
	income2.Amount = 15

	incomes = append(incomes, &income)
	incomes = append(incomes, &income2)
}

func populateExpense() {
	expense := models.Expense{}
	expense.Id = 1
	expense.Amount = 10

	expense2 := models.Expense{}
	expense2.Id = 2
	expense2.Amount = 30

	expenses = append(expenses, &expense)
	expenses = append(expenses, &expense2)
}

func (FakeReportRequest) FindIncome() (models.Incomes, error) {
	return incomes, nil
}

func (FakeReportRequest) FindExpense() (models.Expenses, error) {
	return expenses, nil
}

func (r FakeReportRequest) GetDates() [2]string {

	date := [...]string{r.DateStart, r.DateFinish}

	return date
}

func TestReportService_happyflow(t *testing.T) {

	//Arrange
	expectedReport := models.Report{}

	populateIncome()
	populateExpense()

	expectedReport.DateStart = "2020-12-01"
	expectedReport.DateFinish = "2020-12-02"
	expectedReport.Incomes = incomes
	expectedReport.Expenses = expenses

	expectedReport.TotalExpense = 40
	expectedReport.TotalIncome = 25
	expectedReport.Balance = expectedReport.TotalIncome - expectedReport.TotalExpense

	var r services.ReportInterface = FakeReportRequest{DateStart: "2020-12-01", DateFinish: "2020-12-02"}

	//Act
	report, err := services.GetReport(r)

	//Assert
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}

	if !cmp.Equal(expectedReport, report) {
		t.Errorf("values differents")
	}
}
