package main

import (
	"encoding/json"
	"net/http"
)

type Talk struct {
	Name   string
	Length int
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	talk := Talk{Name: "Quentin", Length: 20}
	payload, err := json.Marshal(talk)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	http.HandleFunc("/", routeHandler)
	http.ListenAndServe(":8080", nil)
}
