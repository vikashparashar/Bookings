package handlers

import (
	"net/http"

	"github.com/vikashparashar/bookings/pkg/config"
	"github.com/vikashparashar/bookings/pkg/render"
)

// Repo is the Repository used by the handlers
var Repo *Repository

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	// sending data to the template
	render.RenderTemplates(w, "home.page.tmpl", nil)
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl", nil)
}

func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "general.page.tmpl", nil)
}
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "major.page.tmpl", nil)
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "contact.page.tmpl", nil)
}
func (m *Repository) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "check.page.tmpl", nil)
}
