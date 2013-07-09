package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

var (
	Address    = flag.String("address", "", "the address to host on")
	Port       = flag.Int("port", 8000, "the port to host on")
	cities = []string{
	"Amsterdam", "San Francisco", "Paris", "New York", "Portland",
}

func main() {
	http.HandleFunc("/", handleIndex)

	address := fmt.Sprintf("%v:%v", *Address, *Port)
	fmt.Printf("Hosting at %v\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func handleIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode(cities)
}
