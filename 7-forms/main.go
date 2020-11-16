package main

import (
	"fmt"
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

		data := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		fmt.Fprint(w, data)
	})

	http.ListenAndServe(":5050", nil)
}
