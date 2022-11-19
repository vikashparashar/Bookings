package main

import (
	"fmt"
	"net/http"

	"github.com/vikashparashar/bookings/cmd/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/general", handlers.General)
	http.HandleFunc("/major", handlers.Major)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/check", handlers.CheckAvailability)

	fmt.Println("\t ->---->>    Starting The Application On Port : 8080    <<----<-")
	http.ListenAndServe(":8080", nil)
}
