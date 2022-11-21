package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vikashparashar/bookings/pkg/config"
	"github.com/vikashparashar/bookings/pkg/handlers"
	"github.com/vikashparashar/bookings/pkg/render"
	"github.com/vikashparashar/bookings/pkg/routes"
)

const (
	portNumber string = ":8080"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("can not create template")
	}

	app.TemplateCache = tc
	app.UseCache = false // in development mode useCache need to be False otherwise in production mode it sets to be True
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/general", handlers.Repo.General)
	// http.HandleFunc("/major", handlers.Repo.Major)
	// http.HandleFunc("/contact", handlers.Repo.Contact)
	// http.HandleFunc("/check", handlers.Repo.CheckAvailability)
	// http.ListenAndServe(portNumber, r)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes.Routes(&app),
	}
	fmt.Println("\t ->---->>    Starting The Application On Port : 8080    <<----<-")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}
