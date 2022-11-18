package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplates(w, "index.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplates(w, "index.html")
}
