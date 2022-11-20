package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vikashparashar/bookings/pkg/config"
)

// stored parsed template into it
// tc is for template cache
var tc = make(map[string]*template.Template)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplates(w http.ResponseWriter, temp string) {
	// get the template cache from the app config
	tc := app.TemplateCache
	// create a template cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatalln("error parsing template files", err)
	// }
	// get requestd templated cache
	t, ok := tc[temp]
	if !ok {
		log.Fatalln("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)

	if err != nil {
		log.Fatalln("could ")
	}
	// rander the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error writing template to browser")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
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
