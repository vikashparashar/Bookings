package handlers

import (
	"net/http"

	"github.com/vikashparashar/bookings/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")
}

func General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "general.page.tmpl")
}
func Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "major.page.tmpl")
}
func Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "contact.page.tmpl")
}
func CheckAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "check.page.tmpl")
}
