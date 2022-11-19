package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// stored parsed template into it
// tc is for template cache
var tc = make(map[string]*template.Template)

/*
// 1st Way For Rendering Template
// this function will read 2 files from disk every single time we run our app and this cause memory cost
func RenderTemplates(w http.ResponseWriter, t string) {
	tmpl, err := template.ParseFiles("./templates/"+t, "./templates/base.layout.tmpl")
	if err != nil {
		log.Fatalln("error parsing template files", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("error parsing template's : ", err)
	}
}
*/

// 2nd Way For Rendering Template
// but now we will create a function which will read file from disk only a single time , so waistage of memory uses

func RenderTemplates(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	// check to see if we already have template into cache

	// if exists use it else create it

	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("Creating The Template .")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have template
		log.Println("Using Cached Template .")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("error parsing template's : ", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s" + t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Fatalln("error parsing template files", err)
	}

	// adding parsed template into cache
	tc[t] = tmpl
	return nil
}
