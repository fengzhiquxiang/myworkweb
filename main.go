package main

import (
	"html/template"
	"log"
	"net/http"
)

type custom struct {
	id          int
	fullname    string
	nickname    string
	address     string
	tradetime   int
}

type product struct {
	name    string
	model    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./js"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	})

	http.HandleFunc("/submit_input_data", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.gohtml", "jquery ajax data back!!")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
