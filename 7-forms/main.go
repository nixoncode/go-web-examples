package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	tmpl := template.Must(template.ParseFiles("contact.html"))

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_ = tmpl.Execute(w, nil)
			return
		}
	})

	http.ListenAndServe(":5050", nil)
}
