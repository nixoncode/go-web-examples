package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)

	// path prefixes
	bookRouter := r.PathPrefix("/books").Subrouter()

	bookRouter.HandleFunc("", allBooks).Methods("GET")
	bookRouter.HandleFunc("/{title}", createBook).Methods("POST")
	bookRouter.HandleFunc("/{title}", readBook).Methods("GET")
	bookRouter.HandleFunc("/{title}/{newTitle}", updateBook).Methods("PUT")
	bookRouter.HandleFunc("/{title}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":5050", r)
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List all the books")
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
