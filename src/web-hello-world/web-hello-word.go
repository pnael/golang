package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	http.ListenAndServe(":8080", nil)
}
