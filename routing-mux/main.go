package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to site. Powered by gorilla mux route handler")
	})

	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello %s", mux.Vars(r)["name"])
	})

	http.ListenAndServe(":5050", r)
}
