package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type User struct {
	Id       int    `json:"id"  gorm:"primary_key"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Users []*User

func (m *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *User) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
