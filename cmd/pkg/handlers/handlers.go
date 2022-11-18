package handlers

import (
	"net/http"

	"github.com/vikashparashar/bookings/cmd/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "index.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "index.html")
}
