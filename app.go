package hello

import (
	"fmt"
	"net/http"
)

func init() {
	r := http.NewServeMux()
	r.HandleFunc("/", handler)
	r.HandleFunc("/ping", pingHandler)

	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong!")
}
