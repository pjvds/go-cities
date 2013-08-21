package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

var (
	Address = flag.String("address", "", "the address to host on")
	Port    = flag.Int("port", 80, "the port to host on")
	cities  = []string{
		"Amsterdam", "San Francisco", "Paris", "New York", "Portland",
	}
)

func main() {
	flag.Parse()
	endpoint := fmt.Sprintf("%v:%v", *Address, *Port)

	http.HandleFunc("/", handleIndex)

	fmt.Printf("Hosting at %v\n", endpoint)
	if err := http.ListenAndServe(endpoint, nil); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func handleIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode(cities)
}
