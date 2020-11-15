package main

import (
	"html/template"
	"net/http"
)

func main() {

	type Todo struct {
		Title string
		Done  bool
	}

	type TodoData struct {
		Title string
		Todos []Todo
	}

	data := TodoData{
		Title: "My Todo List",
		Todos: []Todo{
			{Title: "Do Laundry", Done: false},
			{Title: "Work on Go by examples more", Done: false},
			{Title: "Build a better web", Done: false},
		},
	}

	tmpl := template.Must(template.ParseFiles("views/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)

	})

	http.ListenAndServe(":5050", nil)
}
