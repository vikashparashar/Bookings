package main

import (
	"fmt"
	"net/http"

	"github.com/vikashparashar/bookings/cmd/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/About", handlers.About)
	fmt.Println("\t ->---->>    Starting The Application On Port : 8080    <<----<-")
	http.ListenAndServe(":8080", nil)
}
