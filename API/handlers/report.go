package handlers

import (
	"net/http"

	"github.com/guilhermecoelho/kakeibo/models"
	"github.com/guilhermecoelho/kakeibo/services"
)

var reportReq = models.ReportRequest{}

func GetReport(resp http.ResponseWriter, r *http.Request) {

	errDecode := reportReq.DecodeBody(*r)
	if errDecode != nil {
		http.Error(resp, "Request format error", http.StatusBadRequest)
		return
	}

	var reportInterface services.ReportInterface = services.ReportRequest{
		DateStart:  reportReq.DateStart,
		DateFinish: reportReq.DateFinish,
	}

	report, errGetReport := services.GetReport(reportInterface)
	if errGetReport != nil {
		http.Error(resp, "Internal error", http.StatusInternalServerError)
		return
	}

	err := report.ToJSON(resp)
	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}
