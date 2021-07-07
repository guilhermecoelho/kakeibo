package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Income struct {
	Id         int      `json:"id"  gorm:"primary_key"`
	Name       string   `json:"name"`
	Date       string   `json:"date"`
	CategoryId int      `json:"categoryId"`
	Amount     float32  `json:"amount"`
	Note       string   `json:"note"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

type Incomes []*Income

func (m *Incomes) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Income) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *Income) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
