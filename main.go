package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"encoding/json"
	// "regexp"
    // "io/ioutil"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

type Custom struct {
	Cid         string        `bson:"cid" json:"cid"`
	Fullname    string        `bson:"fullname" json:"fullname"`
	Nickname    string        `bson:"nickname" json:"nickname"`
	Address     string        `bson:"address" json:"address"`
	Phone       string        `bson:"phone" json:"phone"`
	Email       string        `bson:"email" json:"email"`
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
	http.Handle("/jquery-easyui-1.5.1/", http.StripPrefix("/jquery-easyui-1.5.1", http.FileServer(http.Dir("./jquery-easyui-1.5.1"))))

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
        // err = c.Find(bson.M{}).All(&results)
        err = c.Find(nil).Sort("cid").All(&results)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(results)
		// if err = json.NewEncoder(w).Encode(results); err != nil {
		// 	log.Fatal(err)
		// }
		js, err := json.Marshal(results)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}
        // fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
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
        // err = c.Insert(bson.M{"cid":cid, "fullname":fullname,"nickname":nickname,"address":address,"phone":phone,"email":email})
        if err != nil {
                log.Fatal(err)
        }

        //give result to ajax success function receive then it refresh block
        results := make(map[string]int)
    	results["errorMsg"] = 0

		js, err := json.Marshal(results)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.HandleFunc("/update_custom", func(w http.ResponseWriter, r *http.Request) {
		// tpl.ExecuteTemplate(w, "index.gohtml", "jquery ajax data back!!")
		ocid := r.FormValue("ocid")

		fmt.Println(ocid)

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
        err = c.Update(bson.M{"cid":ocid},bson.M{"cid":cid, "fullname":fullname,"nickname":nickname,"address":address,"phone":phone,"email":email})
        if err != nil {
                log.Fatal(err)
        }

        //give result to ajax success function receive then it refresh block
        results := make(map[string]int)
    	results["errorMsg"] = 0

		js, err := json.Marshal(results)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.HandleFunc("/delete_custom", func(w http.ResponseWriter, r *http.Request) {
		cid := r.FormValue("cid")
		fmt.Println(cid) 

		// Connect to our local mongo
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
        err = c.Remove(bson.M{"cid":cid})
        if err != nil {
                log.Fatal(err)
        }

        results := make(map[string]int)
    	results["success"] = 200

		js, err := json.Marshal(results)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		fmt.Println(w)
	})

	http.HandleFunc("/get_search_customs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		var sw string = r.FormValue("search_word")
		fmt.Println(sw) 
		if sw == "" {
			return
		}
		// Connect to our local mongo
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
    	var results []Custom

	    index := mgo.Index{
	      Key: []string{"$text:cid", "$text:fullname", "$text:nickname", "$text:address", "$text:phone", "$text:email"},
	    }

    	err = c.EnsureIndex(index)
        if err != nil {
                log.Fatal(err)
        }
        // sw = "/.*" + sw + ".*/"
        sw = "^" + sw 
		fmt.Println(sw) 
        // err = c.Find(bson.M{"$text": bson.M{"$search": sw}}).Sort("cid").All(&results)
        // err = c.Find(bson.M{"$text": bson.RegEx{regexp.QuoteMeta(sw), ""}}).Sort("cid").All(&results)
        // err = c.Find({$text: { $regex: /2/, $options: 'im' }}).Sort("cid").All(&results)
        err = c.Find(bson.M{"phone": bson.M{"$regex": bson.RegEx{sw, ""}}}).Sort("cid").All(&results)
        // err = c.Find(bson.M{"$text": bson.M{"$search": bson.M{"$regex": bson.RegEx{sw, "m"}}}}).Sort("cid").All(&results)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(results)
		// if err = json.NewEncoder(w).Encode(results); err != nil {
		// 	log.Fatal(err)
		// }
		js, err := json.Marshal(results)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}
        // fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.HandleFunc("/ajax", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		p1 := r.FormValue("name")
		p2 := r.FormValue("city")
		fmt.Println(p1) 
		fmt.Println(p2) 

		results := make(map[string]int)
    	results["success"] = 200

		js, err := json.Marshal(results)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
