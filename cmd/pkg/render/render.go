package render

import (
	"html/template"
	"log"
	"net/http"
)

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
