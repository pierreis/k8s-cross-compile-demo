package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":7777", nil)
}
