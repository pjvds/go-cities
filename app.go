package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", pingHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong!")
}
