package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Logins []*Login

func (m *Logins) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Login) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (mov *Login) DecodeBody(r http.Request) error {

	errorForm := r.ParseForm()
	if errorForm != nil {
		return errors.New(errorForm.Error())
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&mov)
}
