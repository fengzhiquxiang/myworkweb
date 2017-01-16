package main

import (
	// "fmt"
	// "html"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// http.Handle("/foo", fooHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
