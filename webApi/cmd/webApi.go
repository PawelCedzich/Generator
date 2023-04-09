package main

import (
	"html/template"
	"log"
	"strconv"

	"net/http"

	"main.go/webApi/data"
)

type Page struct {
	Title string
	Body  data.BodyData
}

var templates = template.Must(template.ParseFiles("generate.html"))

func loadPage(title string) (*Page, error) {
	return &Page{Title: title}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/generate/"):]
	p, _ := loadPage(title)

	if r.Method == "POST" {
		r.ParseForm()
		p.Body.TableChoice = r.FormValue("options")
		p.Body.AllTable, _ = strconv.ParseFloat(r.FormValue("allTable"), 64)
		p.Body.SingleTable, _ = strconv.ParseFloat(r.FormValue("singleTable"), 64)
	}

	data.DbGenerate(p.Body)

	renderTemplate(w, "generate", p)
}

func main() {
	http.HandleFunc("/generate/", generateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
