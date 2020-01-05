package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

const (
	DBHost  = "127.0.0.1"
	DBUser  = "root"
	DBPass  = "ganesh23"
	DBDbase = "cms"
	PORT    = ":8080"
)

var database *sql.DB

type Page struct {
	Title   string
	Content template.HTML
	Date    string
	GUID    string
}

func (p Page) TruncatedText() string {
	if len(p.Content) > 150 {
		return string(p.Content[:150] + ` ...`)
	}
	return string(p.Content)
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}

	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", pageGUID).
		Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page!")
		return
	}

	t, _ := template.ParseFiles("templates/blog.html")
	err = t.Execute(w, thisPage)
	if err != nil {
		log.Fatal(err)
	}
}

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", 301)
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	Pages := []Page{}
	pageResult, err := database.Query("SELECT page_title,page_content,page_date,page_guid FROM pages ORDER BY ? DESC", "page_date")
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	defer pageResult.Close()

	for pageResult.Next() {
		thisPage := Page{}
		pageResult.Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date, &thisPage.GUID)
		Pages = append(Pages, thisPage)
	}
	t, _ := template.ParseFiles("templates/index.html")
	err = t.Execute(w, Pages)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect!")
		log.Println(err.Error())
	}
	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", ServePage)
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)
	http.Handle("/", routes)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
