package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	cities = []string{
		"Amsterdam", "San Francisco", "Paris", "New York", "Portland",
	}
)

func init() {
	r := http.NewServeMux()
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/ping", pingHandler)

	http.Handle("/", r)
}

func handleIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode(cities)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong!")
}
