package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Για Σας")
}
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Este es el Endpoint")
}
func HandleTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Userslist")
}
func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
	}
	fmt.Fprintf(w, "Payload : %v\n", metadata)
}
func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
