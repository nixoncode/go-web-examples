package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You requested to view %s", r.URL.Path)
	})

	http.ListenAndServe(":5050", nil)
}
