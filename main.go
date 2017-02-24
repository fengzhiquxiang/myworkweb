package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
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

	http.HandleFunc("/add_custom", func(w http.ResponseWriter, r *http.Request) {
		// tpl.ExecuteTemplate(w, "index.gohtml", "jquery ajax data back!!")
		// fmt.Println(r.FormValue("id"))
		fmt.Println(r.FormValue("firstname"))
		fmt.Println(r.FormValue("lastname"))
		fmt.Println(r.FormValue("phone"))
		fmt.Println(r.FormValue("email"))

		// Connect to our local mongo
		mgd, err := mgo.Dial("mongodb://localhost")
		if err != nil {
                panic(err)
        }
        defer mgd.Close()

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
                log.Fatal(err)
        }
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
