package main

import (
	"fmt"
	"net/http"

	"github.com/vikashparashar/bookings/cmd/pkg/config"
	"github.com/vikashparashar/bookings/cmd/pkg/handlers"
)

const (
	portNumber string = ":8080"
)

func main() {
	var app config.AppConfig
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/general", handlers.General)
	http.HandleFunc("/major", handlers.Major)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/check", handlers.CheckAvailability)

	fmt.Println("\t ->---->>    Starting The Application On Port : 8080    <<----<-")
	http.ListenAndServe(portNumber, nil)
}
