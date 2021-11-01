package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type MetData interface {
}
type User struct {
	Dni    string `json:"dni"`
	Name   string `json:"name"`
	Edad   int    `json:"edad"`
	Ciudad string `json:"ciudad"`
	Pais   string `json:"pais"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
