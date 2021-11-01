package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ReportRequest struct {
	DateStart  string `json:"dateStart"`
	DateFinish string `json:"dateFinish"`
}

func (m *ReportRequest) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *ReportRequest) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
