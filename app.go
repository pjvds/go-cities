package app

import (
	"encoding/json"
	"flag"
	"net/http"
)

var (
	Address = flag.String("address", "", "the address to host on")
	Port    = flag.Int("port", 80, "the port to host on")
	cities  = []string{
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
