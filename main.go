package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var envFile = ".env"

func main() {

	http.HandleFunc("/endless_log", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<h1>Hello Oleg!</h1>")

		for {
			log.Print("main | Privet, Oleg!")
			time.Sleep(1 * time.Minute)
			fmt.Fprintf(w, "<h4>Hello Oleg!</h4>")
		}
	})
	http.ListenAndServe(":8080", nil)
}
