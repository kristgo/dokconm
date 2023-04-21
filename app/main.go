package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Subscriber struct {
	ID        int
	Firstname string
	Lastname  string
	Birthyear int
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", subscriberList)
	http.ListenAndServe(":8080", nil)
}

func subscriberList(w http.ResponseWriter, r *http.Request) {
	var age string
	var result []Subscriber
	if age = r.FormValue("subs"); len(age) == 2 {
		result = registerWithLim(age)
	} else if age = r.FormValue("suds"); age == "all" {
		result = registerAll()
	} else {
		result = registerAll()
	}
	templ, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := templ.Execute(w, result); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func registerAll() []Subscriber {
	db, err := sql.Open("mysql", "root:root@tcp(docmysql)/sitedb")
	if err != nil {
		log.Fatal("Open: ", err)
	}
	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping: ", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from sitedb.subscriber")
	if err != nil {
		log.Fatal("Query: ", err)
	}
	defer rows.Close()
	subscribers := []Subscriber{}

	for rows.Next() {
		p := Subscriber{}
		err := rows.Scan(&p.ID, &p.Firstname, &p.Lastname, &p.Birthyear)
		if err != nil {
			log.Fatal("Scan: ", err)
			continue
		}
		subscribers = append(subscribers, p)
	}
	return subscribers
}

func registerWithLim(agelim string) []Subscriber {
	eoutput := make(chan []Subscriber)
	var wg sync.WaitGroup
	wg.Add(1)
	go handleDB(eoutput, &wg, agelim)
	defer close(eoutput)
	wg.Wait()
	return <-eoutput
}

func handleDB(output chan []Subscriber, wg *sync.WaitGroup, agelmt string) {
	agelm, _ := strconv.Atoi(agelmt)
	in := time.Now().Year() - agelm
	var db *sql.DB
	var err error
	db, _ = sql.Open("mysql", "root:root@tcp(docmysql)/sitedb")
	err = db.Ping()
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	p := []Subscriber{}
	rows, err := db.Query("select * from sitedb.subscriber WHERE birthyear < ?", in)
	if err != nil {
		log.Fatal("Query: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		ps := Subscriber{}
		err := rows.Scan(&ps.ID, &ps.Firstname, &ps.Lastname, &ps.Birthyear)
		if err != nil {
			log.Fatal("rows.Scan: ", err)
			continue
		}
		p = append(p, ps)
	}
	wg.Done()
	output <- p
}
