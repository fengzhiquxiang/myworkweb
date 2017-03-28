package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	// "regexp"
    // "io/ioutil"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

type Custom struct {
	Cid         int        `bson:"cid" json:"cid"`
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

	http.HandleFunc("/get_customs_count", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)

		// Connect to our local mongo
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
    	var count int;
        count, err = c.Count()
		fmt.Println("total:"+strconv.Itoa(count))
        if err != nil {
                log.Fatal(err)
        }

        //give result to ajax success function receive then it refresh block
        results := make(map[string]int)
    	results["count"] = count

		js, err := json.Marshal(results)
		if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
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
        fmt.Println("results")
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

	http.HandleFunc("/get_customs_pager", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)

		//pagesz
		pagesz := r.FormValue("pagesz")
		//pageno
		pageno := r.FormValue("pageno")
		fmt.Println("pagesz:"+pagesz)
		fmt.Println("pageno:"+pageno)

		// Connect to our local mongo
		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
	            panic(err)
	    }
	    defer session.Close()
	    c := session.DB("mywebdb").C("customs")
    	var results []Custom
        // err = c.Find(bson.M{}).All(&results)
        err = c.Find(nil).Limit(10).Sort("cid").All(&results)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("results")
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
		scid := r.FormValue("cid")
		cid, err := strconv.Atoi(scid)
		if err != nil {
	            panic(err)
	    }
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
		cid, err := strconv.Atoi(ocid)
		if err != nil {
	            panic(err)
	    }
		fmt.Println(cid)

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
        err = c.Update(bson.M{"cid":cid},bson.M{"cid":cid, "fullname":fullname,"nickname":nickname,"address":address,"phone":phone,"email":email})
        if err != nil {
        		fmt.Println("update customs error!!!!!!!!!")
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
		scid := r.FormValue("cid")
		cid, err := strconv.Atoi(scid)
		if err != nil {
	            panic(err)
	    }
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
                log.Fatal("/delete_custom  ",err)
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
		col := r.FormValue("col")
		val := r.FormValue("val")
		fmt.Println(col,val)
		if val != "" {
			search_customs(w,col,val)
		}
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


func search_customs(w http.ResponseWriter,swname string, swvalue string) {
	fmt.Println(swvalue)
	// Connect to our local mongo
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
            panic(err)
    }
    defer session.Close()
    c := session.DB("mywebdb").C("customs")
	var results []Custom
    //regex
    swvalue = ".*" + swvalue +".*"
	fmt.Println(swvalue) 
    err = c.Find(bson.M{swname: bson.RegEx{swvalue, "im"}}).Sort("cid").All(&results)
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println(results)
	js, err := json.Marshal(results)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
