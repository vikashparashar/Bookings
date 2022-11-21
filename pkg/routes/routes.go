package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vikashparashar/bookings/pkg/config"
	"github.com/vikashparashar/bookings/pkg/handlers"
)

// func routes(app *config.AppConfig) http.Handler {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("welcome"))
// 	})

//		http.ListenAndServe(":3000", r)
//		return r
//	}
func Routes(app *config.AppConfig) http.Handler {

	// 01 using pat routing package

	/* mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/general", http.HandlerFunc(handlers.Repo.General))
	mux.Get("/major", http.HandlerFunc(handlers.Repo.Major))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/check", http.HandlerFunc(handlers.Repo.CheckAvailability))
	*/

	// 02 Using Chi Routing and Middelware Also

	mux := chi.NewRouter()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/general", http.HandlerFunc(handlers.Repo.General))
	mux.Get("/major", http.HandlerFunc(handlers.Repo.Major))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/check", http.HandlerFunc(handlers.Repo.CheckAvailability))
	return mux

}
