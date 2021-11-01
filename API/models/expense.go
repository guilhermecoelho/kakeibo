package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Expense struct {
	Id         int      `json:"id"  gorm:"primary_key"`
	Name       string   `json:"name"`
	Date       string   `json:"date"`
	CategoryId int      `json:"categoryId"`
	Amount     float32  `json:"amount"`
	Note       string   `json:"note"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

type Expenses []*Expense

func (m *Expenses) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Expense) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *Expense) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
