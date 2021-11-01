package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Group struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Groups []*Group

func (m *Groups) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Group) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *Group) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
