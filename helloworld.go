package main

import (
	"fmt"
	"net/http"
)

func nameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", nameHandler)
	http.ListenAndServe(":8080", nil)
}
