package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"encoding/json"
    // "io/ioutil"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

type Custom struct {
	Cid         string        `bson:"cid" json:"cid"`
	Fullname    string        `bson:"fullname" json:"fullname"`
	Nickname    string        `bson:"nickname,omitempty" json:"nickname,omitempty"`
	Address     string        `bson:"address,omitempty" bson:"address,omitempty"`
	Phone       string        `bson:"phone,omitempty" bson:"phone,omitempty"`
	Email       string        `bson:"email,omitempty" bson:"email,omitempty"`
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
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
    	var results []Custom
        err = c.Find(bson.M{}).All(&results)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(results)
		if err = json.NewEncoder(w).Encode(results); err != nil {
			log.Fatal(err)
		}
		// w.WriteHeader(http.StatusOK) // 200
	})

	http.HandleFunc("/add_custom", func(w http.ResponseWriter, r *http.Request) {
		// tpl.ExecuteTemplate(w, "index.gohtml", "jquery ajax data back!!")
		cid := r.FormValue("cid")
		fullname := r.FormValue("fullname")
		nickname := r.FormValue("nickname")
		address := r.FormValue("address")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		fmt.Println(cid)
		fmt.Println(fullname)
		fmt.Println(nickname)
		fmt.Println(address)
		fmt.Println(phone)
		fmt.Println(email)

		// Connect to our local mongo
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
        err = c.Insert(&Custom{cid, fullname,nickname,address,phone,email})
        if err != nil {
                log.Fatal(err)
        }
		// w.WriteHeader(http.StatusOK) // 200
	})

	http.HandleFunc("/delete_custom", func(w http.ResponseWriter, r *http.Request) {
		// tpl.ExecuteTemplate(w, "index.gohtml", "jquery ajax data back!!")
		cid := r.FormValue("cid")
		fullname := r.FormValue("fullname")
		nickname := r.FormValue("nickname")
		address := r.FormValue("address")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		fmt.Println(cid)
		fmt.Println(fullname)
		fmt.Println(nickname)
		fmt.Println(address)
		fmt.Println(phone)
		fmt.Println(email)

		// Connect to our local mongo
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
        err = c.Insert(&Custom{cid, fullname,nickname,address,phone,email})
        if err != nil {
                log.Fatal(err)
        }
		// w.WriteHeader(http.StatusOK) // 200
	})

	http.HandleFunc("/ajax", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		
		// Connect to our local mongo
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
		fmt.Println(c)
        err = c.Insert(&Custom{r.RequestURI, r.RequestURI,"","","",""})
        if err != nil {
                log.Fatal(err)
        }
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
