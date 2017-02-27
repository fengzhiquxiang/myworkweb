package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"encoding/json"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"

)

type custom struct {
	id          int
	fullname    string
	nickname    string
	address     string
	phone       string
	email       string
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

	http.HandleFunc("/get_customs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		// Connect to our local mongo
		mgd, err := mgo.Dial("mongodb://localhost")
		if err != nil {
                panic(err)
        }
        defer mgd.Close()
        var results []custom
        c := mgd.DB("mywebdb").C("customs")
        err = c.Find(nil).All(&results)
        if err != nil {
                log.Fatal(err)
        }
		fmt.Println(results)
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/add_custom", func(w http.ResponseWriter, r *http.Request) {
		// tpl.ExecuteTemplate(w, "index.gohtml", "jquery ajax data back!!")
		// fmt.Println(r.FormValue("id"))
		fmt.Println(r.FormValue("id"))
		fmt.Println(r.FormValue("fullname"))
		fmt.Println(r.FormValue("nickname"))
		fmt.Println(r.FormValue("address"))
		fmt.Println(r.FormValue("phone"))
		fmt.Println(r.FormValue("email"))

		// Connect to our local mongo
		mgd, err := mgo.Dial("mongodb://localhost")
		if err != nil {
                panic(err)
        }
        defer mgd.Close()

        c := mgd.DB("mywebdb").C("customs")
        err = c.Insert(&custom{1, "桂柳化工","桂柳化工","liuzhou","0772-","123@xyz.com"})
        if err != nil {
                log.Fatal(err)
        }
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
