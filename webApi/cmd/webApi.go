package main

import (
	"html/template"
	"log"
	"net/http"

	"main.go/webApi/data"
)

type Page struct {
	Title string
	Body  data.BodyData
}

var templates = template.Must(template.ParseFiles("generator.html", "generate.html"))

func loadPage(title string) (*Page, error) {
	return &Page{Title: title}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func generatorHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/generator/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "generator", p)
}

func generate(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/generate/"):]
	p, _ := loadPage(title)

	p.Body.Val1 = r.FormValue("val1")
	p.Body.Val2 = r.FormValue("val2")
	p.Body.Option = r.FormValue("option")

	data.DbGenerate(p.Body)

	renderTemplate(w, "generator", p)
}

func main() {
	http.HandleFunc("/generator/", generatorHandler)
	http.HandleFunc("/generate/", generate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
