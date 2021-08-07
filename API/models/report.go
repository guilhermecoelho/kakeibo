package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Report struct {
	DateStart    string   `json:"dateStart"`
	DateFinish   string   `json:"dateFinish"`
	TotalIncome  float32  `json:"totalIncome"`
	TotalExpense float32  `json:"totalExpense"`
	Balance      float32  `json:"balance"`
	Expenses     Expenses `json:"Expenses"`
	Incomes      Incomes  `json:"Incomes"`
}

func (m *Report) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *Report) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
