package hello

import (
	"appengine"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("Requested URL: %v", r.URL)

	fmt.Fprint(w, "Hello, world!")
}
