package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// stored parsed template into it
// tc is for template cache
var tc = make(map[string]*template.Template)

// 1st Way For Rendering Template
// this function will read 2 files from disk every single time we run our app and this cause memory cost
func RenderTemplates(w http.ResponseWriter, temp string) {

	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatalln("error parsing template files", err)
	}
	// get requestd templated cache
	t, ok := tc[temp]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	// rander the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	// var myCache = make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all the files name *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all pages ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("failed to parse file : ", ts)
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
