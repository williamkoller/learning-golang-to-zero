package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := user{
			"William",
			"williamkoller@gmail.com",
		}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("listen port 5000")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
