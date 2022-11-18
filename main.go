package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)
	http.ListenAndServe(":8080", nil)
}

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	temp, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Fatalln("error parsing template files", err)
	}
	err = temp.Execute(w, nil)
	if err != nil {
		log.Fatalln("error parsing template's : ", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplates(w, "index.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplates(w, "index.html")
}
