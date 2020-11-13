package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)

	r.HandleFunc("/books/{title}", createBook).Methods("POST")
	r.HandleFunc("/books/{title}", readBook).Methods("GET")
	r.HandleFunc("/books/{title}/{newTitle}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":5050", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to books site")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	fmt.Fprintf(w, "Creating a book with title %s", title)
}

func readBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]

	fmt.Fprintf(w, "Reading a book with title %s", title)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	newTitle := vars["newTitle"]

	fmt.Fprintf(w, "Updating book with title '%s' to new title '%s'", title, newTitle)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	fmt.Fprintf(w, "Deleting a book with title %s", title)
}
