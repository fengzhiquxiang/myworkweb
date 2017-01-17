package main

import (
	"html/template"
	"log"
	"net/http"
)

type custom struct {
	fullname    string
	nickname    string
	traded bool
}

type product struct {
	fullname    string
	nickname    string

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
