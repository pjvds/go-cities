package cities

import (
	"encoding/json"
	"net/http"
)

var (
	cities = []string{
		"Amsterdam", "San Francisco", "Paris", "New York", "Portland",
	}
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode(cities)
}
