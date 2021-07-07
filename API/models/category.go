package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Category struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	GroupId int    `json:"groupId"`
	Group   Group  `gorm:"foreignKey:GroupId"`
}

type Categories []*Category

func (m *Categories) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Category) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *Category) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
