package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vikashparashar/bookings/pkg/config"
	"github.com/vikashparashar/bookings/pkg/handlers"
	"github.com/vikashparashar/bookings/pkg/render"
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

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/general", handlers.General)
	http.HandleFunc("/major", handlers.Major)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/check", handlers.CheckAvailability)

	fmt.Println("\t ->---->>    Starting The Application On Port : 8080    <<----<-")
	http.ListenAndServe(portNumber, nil)
}
