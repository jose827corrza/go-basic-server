package main

import (
	"encoding/json"
	"net/http"
)

//this receives a handler and returns a handler, this
//function enhances the handlers one after other

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Metadata interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
