package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/guilhermecoelho/kakeibo/handlers"
	"github.com/guilhermecoelho/kakeibo/models"
)

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

func TestReport(t *testing.T) {

	reportReq := models.ReportRequest{}
	reportReq.DateStart = "2020-12-01"
	reportReq.DateFinish = "2020-12-02"

	requestBody, _ := json.Marshal(reportReq)

	req, err := http.NewRequest("POST", "/api/report", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	//responseReporter
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetReport)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// expectReport := models.Report{}
	// expectReport.Name = "myname"

	expected := `{"dateStart":"2020-12-01", "dateFinish":"2020-12-02"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
