package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from my handle root")
}

func HandleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from api uri")
}

func HandlePostReq(w http.ResponseWriter, r *http.Request) {
	var metadata Metadata
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload: %v\n", metadata)
}

func HandleRightJson(w http.ResponseWriter, r *http.Request) {
	var user User
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
