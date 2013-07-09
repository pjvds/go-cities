package main

import (
	"encoding/json"
	"net/http"
)

var cities = []string{
	"Amsterdam", "San Francisco", "Paris", "New York", "Portland",
}

func main() {
	http.HandleFunc("/", handleIndex)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		print(err)
	}
}

func handleIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode(cities)
}
